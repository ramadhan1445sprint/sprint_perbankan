package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_segokuning/customErr"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
)

type TransactionRepo interface {
	GetBalanceCurrency(string, string) (*entity.BankAccount, error)
	AddTransaction(string, entity.TransactionPayload) error
	GetBalance(string) ([]entity.ListBalance, error)
}

type transactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) TransactionRepo {
	return &transactionRepo{db}
}

func (r *transactionRepo) GetBalance(userId string) ([]entity.ListBalance, error) {
	var listBalance []entity.ListBalance
	query := "SELECT currency, total_balance FROM bank_accounts WHERE user_id = $1 ORDER BY total_balance DESC"

	err := r.db.Select(&listBalance, query, userId)
	if err != nil {
		return nil, err
	}

	return listBalance, nil
}

func (r *transactionRepo) GetBalanceCurrency(currency string, userId string) (*entity.BankAccount, error) {
	var bankData entity.BankAccount
	query := "SELECT id, currency, total_balance FROM bank_accounts WHERE user_id = $1 and currency = $2"

	err := r.db.Get(&bankData, query, userId, currency)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		}
		return nil, customErr.NewBadRequestError("balance is not enough")
	}

	return &bankData, nil
}

func (r *transactionRepo) AddTransaction(userId string, payload entity.TransactionPayload) (err error) {
	tx := r.db.MustBegin()

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()

			if rollbackErr != nil {
				err = rollbackErr
			}
		}
	}()

	// add transaction
	statement := "INSERT INTO transactions (user_id, account_name, account_number, currency, balance) VALUES ($1, $2, $3, $4, $5)"

	_, err = tx.Exec(statement, userId, payload.RecipientBankName, payload.RecipientBankAccountNumber, payload.FromCurrency, payload.Balances*(-1))
	if err != nil {
		return err
	}

	// increment total_balances
	_, err = tx.Exec("UPDATE bank_accounts SET total_balance = total_balance + $1 WHERE user_id = $2 and currency = $3",
		payload.Balances*(-1),
		userId,
		payload.FromCurrency,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
