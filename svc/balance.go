package svc

import (
	"github.com/ramadhan1445sprint/sprint_segokuning/customErr"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
	"github.com/ramadhan1445sprint/sprint_segokuning/repo"
)

type BalanceSvc interface {
	AddBankAccountBalance(req entity.AddBalanceRequest) *customErr.CustomError
	GetBalanceHistory(userId string, filter entity.BalanceHistoryMeta) (*entity.BalanceHistoryDataResponse, *customErr.CustomError) 
}

type balanceSvc struct {
	repo repo.BalanceRepo
}

func NewBalanceSvc(r repo.BalanceRepo) BalanceSvc {
	return &balanceSvc{
		repo: r,
	}
}

func (s *balanceSvc) AddBankAccountBalance(req entity.AddBalanceRequest) *customErr.CustomError {

	bank := entity.BankAccountBalance{
		UserID: req.UserID,
		Currency: req.Currency,
		TotalBalance: req.Balance,
	}

	transaction := entity.BalanceTransaction{
		UserID: req.UserID,
		AccountNumber: req.AccountNumber,
		BankName: req.BankName,
		Currency: req.Currency,
		ImageUrl: req.TransferProofImg,
		Balance: req.Balance,
	}

	if err := s.repo.AddBankAccountBalance(&bank, &transaction); err != nil {
		custErr := customErr.NewInternalServerError(err.Error())
		return &custErr
	}

	return nil
}

func (s *balanceSvc) GetBalanceHistory(userId string, filter entity.BalanceHistoryMeta) (*entity.BalanceHistoryDataResponse, *customErr.CustomError) {

	resp, err := s.repo.GetBalanceHistory(userId, filter)

	if err != nil {
		custErr := customErr.NewInternalServerError(err.Error())
		return nil, &custErr
	}

	meta := entity.BalanceHistoryMeta{
		Limit: filter.Limit,
		Offset: filter.Offset,
		Total: len(resp),
	}

	balanceResp := entity.BalanceHistoryDataResponse{
		Message: "success",
		Data: resp,
		Meta: meta,
	}

	return &balanceResp, nil
}
