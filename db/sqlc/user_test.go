package db

import (
	"context"
	"github.com/stefan-vl/my-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{Username: util.RandomOwner(), HashedPassword: "secret", FullName: util.RandomOwner(), Email: util.RandomEmail()}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoErrorf(t, err, "", nil)
	require.NotEmptyf(t, user, "", nil)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZerof(t, user.CreatedAt, "", nil)

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoErrorf(t, err, "", nil)
	require.NotEmptyf(t, user, "", nil)
	require.Equal(t, user1.Username, user.Username)
	require.Equal(t, user1.HashedPassword, user.HashedPassword)
	require.Equal(t, user1.FullName, user.FullName)
	require.Equal(t, user1.Email, user.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user.CreatedAt, time.Second)
}
