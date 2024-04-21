package handlers

import (
	"encoding/json"
	modelapp "goFinalTask/modelAPP"
	modeldb "goFinalTask/modelDB"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	base := vars["base"]

	switch r.Method {
	case "GET":
		log.Println(r.Method)
		log.Printf("Getting transaction...")
		if id != "" {
			if base != "" {
				log.Println(".. by currency")
				getTransactionsByCurrency(w, r)
				return
			} else {
				log.Println("... by ID")
				getTransaction(w, r)
				return
			}

		} else {
			log.Println("all transactions")
			getAllTransactions(w, r)
		}
	case "POST":
		createTransaction(w, r)
	case "PUT":
		updateTransaction(w, r)
	case "DELETE":
		deleteTransaction(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}

func getTransaction(w http.ResponseWriter, r *http.Request) {

	var transaction modelapp.Transactions
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	result := modeldb.DB.Find(&transaction, id)
	if result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

func getTransactionsByCurrency(w http.ResponseWriter, r *http.Request) {
	var exchangeResult modelapp.ExchangeResult
	var transaction modelapp.Transactions
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	base, _ := vars["base"]
	log.Println(base, id)
	result := modeldb.DB.Find(&transaction, id)
	if result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}
	rate, err := GetLastWithBase(base, transaction.Currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	exchangeResult.Currency = base
	if rate > 0 {
		exchangeResult.Amount = transaction.Amount / rate
	} else {
		exchangeResult.Amount = 0
	}
	exchangeResult.Amount = transaction.Amount / rate
	exchangeResult.Date = time.Now()
	exchangeResult.Rate = rate
	exchangeResult.Transactions = transaction //append(exchangeResult.Transactions, transaction)
	//log.Println(exchangeResult)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchangeResult)
	return
}

func getAllTransactions(w http.ResponseWriter, r *http.Request) {
	// Get all transactions from the database
	var transactions []modelapp.Transactions
	result := modeldb.DB.Find(&transactions)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)

}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction modelapp.Transactions
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Commission(&transaction)

	transaction.CreateDate = time.Now()

	result := modeldb.DB.Create(&transaction)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func updateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction modelapp.Transactions
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	result := modeldb.DB.Find(&modelapp.Transactions{}, id)
	if result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transaction.ID = uint(id)
	Commission(&transaction)
	transaction.CreateDate = time.Now().UTC()
	result = modeldb.DB.Save(&transaction)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
	log.Println("Transaction updated successfully!")
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//var transaction modelapp.Transactions

	result := modeldb.DB.Find(&modelapp.Transactions{}, id)
	if result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}
	result = modeldb.DB.Delete(&modelapp.Transactions{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func Commission(c *modelapp.Transactions) {

	switch {
	case c.Transaction == "transfer", c.Transaction == "перевод" && c.Currency != "RUB":
		c.AmountPaid = c.Amount * 0.02
	case c.Transaction == "transfer", c.Transaction == "перевод" && c.Currency == "RUB":
		c.AmountPaid = c.Amount * 0.05
	default:
		c.AmountPaid = 0
	}

}
