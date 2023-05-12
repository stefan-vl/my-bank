package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transferParams := CreateTransferParams{FromAccountID: account2.ID, ToAccountID: account1.ID, Amount: rand.Int63()}

	transfer1, err := testQueries.CreateTransfer(context.Background(), transferParams)
	require.NoError(t, err, "", nil)
	require.NotEmptyf(t, transfer1, "", nil)
	require.NotZerof(t, transfer1.ID, "", nil)
	require.NotZerof(t, transfer1.CreatedAt, "", nil)
	require.Equal(t, transfer1.Amount, transferParams.Amount)
	require.Equal(t, transfer1.ToAccountID, transferParams.ToAccountID)
	require.Equal(t, transfer1.FromAccountID, transferParams.FromAccountID)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transferParams := CreateTransferParams{FromAccountID: account2.ID, ToAccountID: account1.ID, Amount: rand.Int63()}

	createdTransfer, _ := testQueries.CreateTransfer(context.Background(), transferParams)

	transfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)

	require.NoError(t, err, "", nil)
	require.NotEmptyf(t, createdTransfer, "", nil)
	require.Equal(t, createdTransfer.ID, transfer.ID)
	require.WithinDuration(t, createdTransfer.CreatedAt, transfer.CreatedAt, time.Second)
	require.Equal(t, createdTransfer.Amount, transfer.Amount)
	require.Equal(t, createdTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, createdTransfer.FromAccountID, transfer.FromAccountID)
}
