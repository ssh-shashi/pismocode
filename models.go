package main

type Account struct {
    AccountID      int    `json:"account_id"`
    DocumentNumber string `json:"document_number"`
}

type OperationType struct {
    OperationTypeID int    `json:"operation_type_id"`
    Description     string `json:"description"`
}

type Transaction struct {
    TransactionID   int     `json:"transaction_id"`
    AccountID       int     `json:"account_id"`
    OperationTypeID int     `json:"operation_type_id"`
    Amount          float64 `json:"amount"`
    EventDate       string  `json:"event_date"`
}
