package db

import (
	"context"
	"github.com/jammpu/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		1,
		util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry.ID)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.Amount, entry2.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		1,
		10,
		0,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 10)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
