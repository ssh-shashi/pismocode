package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
        return
    }

    var account Account
    err := json.NewDecoder(r.Body).Decode(&account)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    account = createAccount(account)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(account)
}

func getAccountHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
        return
    }

    idStr := strings.TrimPrefix(r.URL.Path, "/accounts/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid account ID.", http.StatusBadRequest)
        return
    }

    account, err := getAccount(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(account)
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
        return
    }

    var transaction Transaction
    err := json.NewDecoder(r.Body).Decode(&transaction)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    transaction = createTransaction(transaction)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(transaction)
}
