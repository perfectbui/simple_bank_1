package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(10)
	testcaseQty := 5
	errs := make(chan error)
	results := make(chan TransferTxResult)
	// for i := 0; i < testcaseQty; i++ {
	// 	result, err := store.TransferTx(context.Background(), CreateTransferParams{FromAccountID: account1.ID, ToAccountID: account2.ID, Amount: amount})

	// 	require.NoError(t, err)
	// 	require.NotEmpty(t, result)
	// 	require.Equal(t, account1.ID, result.Transfers.FromAccountID)
	// 	require.Equal(t, account2.ID, result.Transfers.ToAccountID)
	// 	require.Equal(t, amount, result.Transfers.Amount)

	// 	transfer, err := store.GetTranser(context.Background(), result.Transfers.ID)
	// 	require.NoError(t, err)
	// 	require.NotEmpty(t, transfer)

	// 	fromEntry, err := store.GetEntry(context.Background(), result.FromEntry.ID)
	// 	require.NoError(t, err)
	// 	require.NotEmpty(t, fromEntry)
	// 	require.Equal(t, account1.ID, fromEntry.AccountID)
	// 	require.Equal(t, -amount, fromEntry.Amount)

	// 	toEntry, err := store.GetEntry(context.Background(), result.ToEntry.ID)
	// 	require.NoError(t, err)
	// 	require.NotEmpty(t, toEntry)
	// 	require.Equal(t, account2.ID, toEntry.AccountID)
	// 	require.Equal(t, amount, toEntry.Amount)

	// 	// check accounts
	// 	// result.FromAccount.Balance
	// }

	for i := 0; i < testcaseQty; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), CreateTransferParams{FromAccountID: account1.ID, ToAccountID: account2.ID, Amount: amount})
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < testcaseQty; i++ {
		go func() {
			err := <-errs
			result := <-results
			require.NoError(t, err)
			require.NotEmpty(t, result)
			require.Equal(t, account1.ID, result.Transfers.FromAccountID)
			fmt.Println(account1.ID, result.Transfers.FromAccountID)
			require.Equal(t, account2.ID, result.Transfers.ToAccountID)
			require.Equal(t, amount, result.Transfers.Amount)

			transfer, err := store.GetTranser(context.Background(), result.Transfers.ID)
			require.NoError(t, err)
			require.NotEmpty(t, transfer)

			fromEntry, err := store.GetEntry(context.Background(), result.FromEntry.ID)
			require.NoError(t, err)
			require.NotEmpty(t, fromEntry)
			require.Equal(t, account1.ID, fromEntry.AccountID)
			require.Equal(t, -amount, fromEntry.Amount)

			toEntry, err := store.GetEntry(context.Background(), result.ToEntry.ID)
			require.NoError(t, err)
			require.NotEmpty(t, toEntry)
			require.Equal(t, account2.ID, toEntry.AccountID)
			require.Equal(t, amount, toEntry.Amount)
		}()
	}

}
