package db_test

import (
	"context"
	"testing"

	db "github.com/caiofernandes00/simple_bank/src/infrastructure/db/sqlc"
	"github.com/stretchr/testify/require"
)

func Test_CreateAccount(t *testing.T) {
	arg := db.CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
