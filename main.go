package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/accounts", createAccountHandler)
    http.HandleFunc("/accounts/", getAccountHandler)
    http.HandleFunc("/transactions", createTransactionHandler)

    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
