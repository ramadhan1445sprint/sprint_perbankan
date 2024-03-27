package svc

import (
	"github.com/ramadhan1445sprint/sprint_segokuning/customErr"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
	"github.com/ramadhan1445sprint/sprint_segokuning/repo"
)

type TransactionSvc interface {
	AddTransaction(userId string, payload entity.TransactionPayload) error
	GetBalance(userId string) ([]entity.ListBalance, error)
}

type transactionSvc struct {
	repo repo.TransactionRepo
}

func NewTransactionSvc(repo repo.TransactionRepo) TransactionSvc {
	return &transactionSvc{repo}
}

func (s *transactionSvc) AddTransaction(userId string, payload entity.TransactionPayload) error {
	if err := payload.Validate(); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	// check bank account exist or not, if not insert bank account
	bankDetail, err := s.repo.GetBalanceCurrency(payload.FromCurrency, userId)
	if err != nil {
		return err
	}

	if bankDetail != nil {
		if bankDetail.TotalBalance < payload.Balances {
			return customErr.NewBadRequestError("balance is not enough")
		}
	}

	// add transaction
	err = s.repo.AddTransaction(userId, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *transactionSvc) GetBalance(userId string) ([]entity.ListBalance, error) {
	listBalance, err := s.repo.GetBalance(userId)
	if err != nil {
		return nil, err
	}

	return listBalance, nil
}
