package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
)

type BalanceRepo interface {
	AddBankAccountBalance(bank *entity.BankAccountBalance, transaction *entity.BalanceTransaction) error
	GetBalanceHistory(userId string, filter entity.BalanceHistoryMeta) ([]entity.BalanceHistoryData, error)
}

type balanceRepo struct {
	db *sqlx.DB
}

func NewBalanceRepo(db *sqlx.DB) BalanceRepo {
	return &balanceRepo{
		db: db,
	}
}

func (r *balanceRepo) AddBankAccountBalance(bank *entity.BankAccountBalance, transaction *entity.BalanceTransaction) error {
	tx, err := r.db.BeginTx(context.Background(), nil)

	defer tx.Rollback()

	var exist int
	var query string

	if err = tx.QueryRow("SELECT count(*) from bank_accounts where user_id = $1 and currency = $2", bank.UserID, bank.Currency).Scan(&exist); err != nil {
		return err
	}

	if exist > 0 {
		query = "UPDATE bank_accounts SET total_balance = total_balance + $1 WHERE user_id = $2 and currency = $3"
	} else {
		query = "INSERT INTO bank_accounts (total_balance, user_id, currency) VALUES ($1, $2, $3)"
	}

	_, err = tx.Exec(query, bank.TotalBalance, bank.UserID, bank.Currency)

	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, account_name, account_number, currency, balance, image_url) VALUES ($1, $2, $3, $4, $5, $6)", transaction.UserID, transaction.BankName, transaction.AccountNumber, transaction.Currency, transaction.Balance, transaction.ImageUrl)

	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *balanceRepo) GetBalanceHistory(userId string, filter entity.BalanceHistoryMeta) ([]entity.BalanceHistoryData, error) {
	query := `
		SELECT 
			id,
			balance,
			currency,
			image_url,
			created_at, 
			account_number,
			account_name
			from transactions 
			WHERE user_id = $1
			ORDER BY created_at DESC limit %d offset %d
	`
	query = fmt.Sprintf(query, filter.Limit, filter.Offset)

	rows, err := r.db.Query(query, userId)

	if err != nil {
		return nil, err
	}

	balanceHistory := []entity.BalanceHistoryData{}

	for rows.Next() {
		var tempRawData entity.BalanceHistoryRawData
		var tempSourceData entity.BalanceHistorySourceData
		var tempData entity.BalanceHistoryData

		if err := rows.Scan(&tempRawData.ID, &tempRawData.Balance, &tempRawData.Currency, &tempRawData.TransferProofImg, &tempRawData.CreatedAt, &tempRawData.AccountNumber, &tempRawData.BankName); err != nil {
			return nil, err
		}

		tempSourceData.AccountNumber = tempRawData.AccountNumber
		tempSourceData.BankName = tempRawData.BankName

		if tempRawData.Balance < 0 {
			tempRawData.Balance = tempRawData.Balance * (-1)
		}

		tempData.ID = tempRawData.ID
		if tempRawData.Balance < 0 {
			tempRawData.Balance *= -1
		}
		tempData.Balance = tempRawData.Balance
		tempData.Currency = tempRawData.Currency
		tempData.TransferProofImg = tempRawData.TransferProofImg
		tempData.CreatedAt = int(tempRawData.CreatedAt.UnixNano()) / int(time.Millisecond)
		tempData.Source = tempSourceData
		balanceHistory = append(balanceHistory, tempData)
	}

	return balanceHistory, nil
}
