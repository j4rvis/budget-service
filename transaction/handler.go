package transaction

import "net/http"

// Transaction ...
type Transaction struct{}

// NewHandler ...
func NewHandler() *Transaction {
	transaction := Transaction{}

	return &transaction
}

// GetTransactions ...
func (t *Transaction) GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Some Transactions"))
}
