package entity

import "time"

type BankAccountBalance struct {
	UserID       string `json:"userId"`
	Currency     string `json:"currency"`
	TotalBalance string `json:"totalBalance"`
}

type BalanceTransaction struct {
	UserID        string `json:"userId"`
	AccountNumber string `json:"senderBankAccountNumber"`
	BankName      string `json:"senderBankName"`
	Currency      string `json:"currency"`
	ImageUrl      string `json:"imageUrl"`
	Balance       int    `json:"balance"`
}

type BalanceHistoryData struct {
	ID               string                   `json:"transactionId"`
	Balance          int                      `json:"balance"`
	Currency         string                   `json:"currency"`
	TransferProofImg string                   `json:"transferProofImg"`
	CreatedAt        int                      `json:"createdAt"`
	Source           BalanceHistorySourceData `json:"source"`
}

type BalanceHistoryRawData struct {
	ID               string    `json:"transactionId"`
	Balance          int       `json:"balance"`
	Currency         string    `json:"currency"`
	TransferProofImg string    `json:"transferProofImg"`
	CreatedAt        time.Time `json:"createdAt"`
	AccountNumber    string    `json:"bankAccountNumber"`
	BankName         string    `json:"bankName"`
}

type BalanceHistorySourceData struct {
	AccountNumber string `json:"bankAccountNumber"`
	BankName      string `json:"bankName"`
}

type BalanceHistoryMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type BalanceHistoryDataResponse struct {
	Message string               `json:"message"`
	Data    []BalanceHistoryData `json:"data"`
	Meta    BalanceHistoryMeta   `json:"meta"`
}
