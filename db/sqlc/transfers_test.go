package db

import (
	"context"
	"testing"
	"time"

	"github.com/snirkop89/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	expTranfer := createRandomTransfer(t, account1, account2)

	gotTransfer, err := testQueries.GetTransfer(context.Background(), expTranfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotTransfer)

	require.Equal(t, expTranfer.ID, gotTransfer.ID)
	require.Equal(t, expTranfer.Amount, gotTransfer.Amount)
	require.Equal(t, expTranfer.FromAccountID, gotTransfer.FromAccountID)
	require.Equal(t, expTranfer.ToAccountID, gotTransfer.ToAccountID)
	require.WithinDuration(t, expTranfer.CreatedAt, gotTransfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID && transfer.ToAccountID == account2.ID)
	}
}
