package repo

import (
	"testing"

	"github.com/ramadhan1445sprint/sprint_segokuning/config"
	"github.com/ramadhan1445sprint/sprint_segokuning/db"
	"github.com/ramadhan1445sprint/sprint_segokuning/entity"
)

func TestAndBankAccountBalance(t *testing.T) {
	config.LoadConfig("../.env")

	db, err := db.NewDatabase()
	if err != nil {
		t.Fatalf("failed to create a database connection: %v", err)
	}

	balanceRepo := NewBalanceRepo(db)

	testCases := []struct {
		name        string
		bank       entity.BankAccountBalance
		transaction entity.BalanceTransaction
		errExpected bool
	}{
		{"Test add new bank balance", entity.BankAccountBalance{UserID: "979b12b3-e4da-479e-8418-31743a1c63d1", Currency: "IDR", TotalBalance: 5000},entity.BalanceTransaction{UserID: "979b12b3-e4da-479e-8418-31743a1c63d1", AccountNumber: "120210", BankName: "BCA", Currency: "IDR", ImageUrl: "dwdw", Balance: 5000}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := balanceRepo.AddBankAccountBalance(&tc.bank, &tc.transaction)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}

func TestGetBalanceHistory(t *testing.T) {
	config.LoadConfig("../.env")

	db, err := db.NewDatabase()
	if err != nil {
		t.Fatalf("failed to create a database connection: %v", err)
	}

	balanceRepo := NewBalanceRepo(db)

	testCases := []struct {
		name        string
		userId      string
		filter      entity.BalanceHistoryMeta
		errExpected bool
	}{
		{"Test add new bank balance", "979b12b3-e4da-479e-8418-31743a1c63d1", entity.BalanceHistoryMeta{Limit: 2, Offset: 0}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := balanceRepo.GetBalanceHistory(tc.userId, tc.filter)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}