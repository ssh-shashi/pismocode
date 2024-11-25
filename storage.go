package main

import (
    "errors"
    "time"
)

var accounts = make(map[int]Account)
var transactions = make(map[int]Transaction)
var accountIDCounter = 1
var transactionIDCounter = 1

func createAccount(account Account) Account {
    account.AccountID = accountIDCounter
    accounts[accountIDCounter] = account
    accountIDCounter++
    return account
}

func getAccount(id int) (Account, error) {
    account, exists := accounts[id]
    if !exists {
        return Account{}, errors.New("account not found")
    }
    return account, nil
}

func createTransaction(transaction Transaction) Transaction {
    transaction.TransactionID = transactionIDCounter
    transaction.EventDate = time.Now().Format(time.RFC3339)
    transactions[transactionIDCounter] = transaction
    transactionIDCounter++
    return transaction
}
