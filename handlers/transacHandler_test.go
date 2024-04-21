package handlers

import (
	modelapp "goFinalTask/modelAPP"
	"testing"
	"time"
)

func TestComissions(t *testing.T) {
	testCreateDate := time.Now().UTC()
	testTransactions := []modelapp.Transactions{
		{ID: 1, UserID: 1, Amount: 100.00, Currency: "USD", Transaction: "transfer", Category: "SBP", CreateDate: testCreateDate, Description: "lorem ipsum 1", AmountPaid: 2.00},
		{ID: 2, UserID: 1, Amount: 200.00, Currency: "EUR", Transaction: "transfer", Category: "SBP", CreateDate: testCreateDate, Description: "lorem ipsum 2", AmountPaid: 4.00},
		{ID: 3, UserID: 1, Amount: 1000.00, Currency: "RUB", Transaction: "перевод", Category: "SBP", CreateDate: testCreateDate, Description: "lorem ipsum 3", AmountPaid: 50.0},
		{ID: 4, UserID: 3, Amount: 1e20, Currency: "GBP", Transaction: "transfer", Category: "SBP", CreateDate: testCreateDate, Description: "lorem ipsum 4", AmountPaid: 2e18},
	}
	for _, tt := range testTransactions {
		transactionDate := tt.CreateDate
		commission := tt.AmountPaid
		Commission(&tt)
		tt.CreateDate = time.Now().UTC()
		if commission != tt.AmountPaid || transactionDate.After(tt.CreateDate) {
			t.Errorf("Commission Amount(%v) = %v, want %v", tt, tt.AmountPaid, commission)
			t.Errorf("Transaction Date(%v) = %v, want %v", tt, tt.CreateDate, transactionDate)
		}
	}

}
