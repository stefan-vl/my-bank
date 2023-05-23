package db

import (
	"context"
	"github.com/stefan-vl/my-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{Owner: user.Username, Balance: util.RandomMoney(), Currency: util.RandomCurrency()}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoErrorf(t, err, "", nil)
	require.NotEmptyf(t, account, "", nil)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZerof(t, account.ID, "", nil)
	require.NotZerof(t, account.CreatedAt, "", nil)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoErrorf(t, err, "", nil)
	require.NotEmptyf(t, account, "", nil)
	require.Equal(t, account1.Owner, account.Owner)
	require.Equal(t, account1.Balance, account.Balance)
	require.Equal(t, account1.Currency, account.Currency)
	require.Equal(t, account1.ID, account.ID)
	require.WithinDuration(t, account1.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	params := UpdateAccountParams{ID: account.ID, Balance: util.RandomMoney()}

	account1, err := testQueries.UpdateAccount(context.Background(), params)
	require.NoErrorf(t, err, "", nil)
	require.NotEmptyf(t, account1, "", nil)
	require.Equal(t, account1.Owner, account.Owner)
	require.Equal(t, account1.Balance, params.Balance)
	require.Equal(t, account1.Currency, account.Currency)
	require.Equal(t, account1.ID, account.ID)
	require.WithinDuration(t, account1.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err, "", nil)

	entry, err := testQueries.GetEntry(context.Background(), account.ID)
	require.Error(t, err, "", nil)
	require.Empty(t, entry, "", nil)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}
	params := ListAccountsParams{Owner: lastAccount.Owner, Limit: 5, Offset: 0}

	accounts, err := testQueries.ListAccounts(context.Background(), params)
	require.NoError(t, err, "", nil)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmptyf(t, account, "", nil)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
