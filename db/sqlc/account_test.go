package db

import (
	"LearningTransaction/simplebank_1/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount1() Accounts {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, _ := testQueries.CreateAccount(context.Background(), arg)

	return account
}

func createRandomAccount(t *testing.T) Accounts {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	for i := 0; i < 3; i++ {
		newAccount := createRandomAccount(t)
		existedAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)
		require.NoError(t, err)
		require.NotEmpty(t, existedAccount)
		require.Equal(t, newAccount.Owner, existedAccount.Owner)
		require.Equal(t, newAccount.Balance, existedAccount.Balance)
		require.Equal(t, newAccount.Currency, existedAccount.Currency)
	}
}

func TestUpdateAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	newBalance := utils.RandomMoney()
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: newBalance,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, newBalance, updatedAccount.Balance)
}

func TestDeleteAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)
}

func TestGetAccountList(t *testing.T) {
	newAccount1 := createRandomAccount(t)
	newAccount2 := createRandomAccount(t)
	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{Offset: 0, Limit: 2})
	require.NoError(t, err)
	require.NotEmpty(t, newAccount1)
	require.Equal(t, newAccount1.Owner, accounts[1].Owner)
	require.Equal(t, newAccount1.Balance, accounts[1].Balance)
	require.Equal(t, newAccount1.Currency, accounts[1].Currency)
	require.NotEmpty(t, newAccount2)
	require.Equal(t, newAccount2.Owner, accounts[0].Owner)
	require.Equal(t, newAccount2.Balance, accounts[0].Balance)
	require.Equal(t, newAccount2.Currency, accounts[0].Currency)
}
