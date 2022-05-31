package db_test

import (
	"context"
	"testing"

	db "github.com/caiofernandes00/simple_bank/src/infrastructure/db/sqlc"
	"github.com/stretchr/testify/require"
)

func Test_TransferTx(t *testing.T) {
	store := db.NewStore(testDB)

	account1 := createAccount(t)
	account2 := createAccount(t)

	// run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan db.TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), db.TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, account1.ID, transfer.FromAccountID)
		require.NotEmpty(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		queryTransfer, err := store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)
		require.Equal(t, transfer, queryTransfer)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		queryFromEntry, err := store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)
		require.Equal(t, fromEntry, queryFromEntry)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		queryToEntry, err := store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
		require.Equal(t, toEntry, queryToEntry)
	}
}
