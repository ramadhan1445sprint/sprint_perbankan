package entity

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TransactionPayload struct {
	RecipientBankAccountNumber string `json:"recipientBankAccountNumber"`
	RecipientBankName          string `json:"RecipientBankName"`
	FromCurrency               string `json:"fromCurrency"`
	Balances                   int    `json:"balances"`
}

type BankAccount struct {
	Id           string `db:"id"`
	Currency     string `db:"currency"`
	TotalBalance int    `db:"total_balance"`
}

type ListBalance struct {
	Currency     string `db:"currency"`
	TotalBalance int    `db:"total_balance"`
}

func (p *TransactionPayload) Validate() error {
	err := validation.ValidateStruct(p,
		validation.Field(&p.RecipientBankAccountNumber,
			validation.Required.Error("bank account number is required"),
			validation.Length(5, 30).Error("bank account number must be between 5 and 30 characters"),
		),
		validation.Field(&p.RecipientBankName,
			validation.Required.Error("bank name is required"),
			validation.Length(5, 30).Error("bank name must be between 5 and 30 characters"),
		),
		validation.Field(&p.FromCurrency,
			validation.Required.Error("Fromcurrency is required"),
			validation.By(validateCurrency),
		),
		validation.Field(&p.Balances,
			validation.Required.Error("balances is required"),
		),
	)

	return err
}

func validateCurrency(value any) error {
	currency, ok := value.(string)
	if !ok {
		return errors.New("parse error")
	}

	pattern := "^[A-Z]{3}$"
	rgx := regexp.MustCompile(pattern)
	if !rgx.MatchString(currency) {
		return errors.New("invalid currency format")
	}

	return nil
}
