package tests

import (
	"testing"

	"github.com/ssh-shashi/pismocode/models"
	"github.com/ssh-shashi/pismocode/storage"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	account := models.Account{DocumentNumber: "12345678900"}
	createdAccount := storage.CreateAccount(account)

	assert.Equal(t, "12345678900", createdAccount.DocumentNumber)
	assert.NotZero(t, createdAccount.AccountID)
}

func TestGetAccount(t *testing.T) {
	account := models.Account{DocumentNumber: "12345678900"}
	createdAccount := storage.CreateAccount(account)

	fetchedAccount, err := storage.GetAccount(createdAccount.AccountID)
	assert.NoError(t, err)
	assert.Equal(t, createdAccount, fetchedAccount)
}
