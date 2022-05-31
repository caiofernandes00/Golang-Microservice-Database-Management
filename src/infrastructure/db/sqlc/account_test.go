package db_test

import (
	"context"
	"testing"

	db "github.com/caiofernandes00/simple_bank/src/infrastructure/db/sqlc"
	utils_test "github.com/caiofernandes00/simple_bank/test/utils"
	"github.com/stretchr/testify/require"
)

func createAccount(t *testing.T) db.Account {
	args := db.CreateAccountParams{
		Owner:    utils_test.RandomOwner(),
		Balance:  utils_test.RandomMoney(),
		Currency: utils_test.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func Test_CreateAccount(t *testing.T) {
	createAccount(t)
}

func Test_ListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createAccount(t)
	}

	args := db.ListAccountsParams{
		Limit:  2,
		Offset: 2,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 2)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
