package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/ssh-shashi/pismocode/models"
)

var (
	accounts     = make(map[int]models.Account)
	transactions = make(map[int]models.Transaction)
	accountID    = 1
	transactionID = 1
	mu           sync.Mutex
)

func CreateAccount(account models.Account) models.Account {
	mu.Lock()
	defer mu.Unlock()
	account.AccountID = accountID
	accounts[accountID] = account
	accountID++
	return account
}

func GetAccount(accountID int) (models.Account, error) {
	mu.Lock()
	defer mu.Unlock()
	account, exists := accounts[accountID]
	if !exists {
		return models.Account{}, errors.New("account not found")
	}
	return account, nil
}

func CreateTransaction(transaction models.Transaction) models.Transaction {
	mu.Lock()
	defer mu.Unlock()
	transaction.TransactionID = transactionID
	transaction.EventDate = time.Now()
	transactions[transactionID] = transaction
	transactionID++
	return transaction
}
