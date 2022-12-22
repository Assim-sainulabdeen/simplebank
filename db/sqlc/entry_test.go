package db

import (
	"context"
	"reflect"
	"testing"

	"github.com/Assim-sainulabdeen/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account Account) Entry {

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	CreateRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account1 := CreateRandomAccount(t)
	entry1 := CreateRandomEntry(t, account1)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	// require.NotEmpty(t, entry2)
	// require.Equal(t, entry1.AccountID, entry2.AccountID)
	// require.Equal(t, entry1.Amount, entry2.Amount)
	// require.Equal(t, entry1.ID, entry2.ID)
	// require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
	isEqual := reflect.DeepEqual(entry1, entry2)
	require.True(t, isEqual)
}

func TestListEntries(t *testing.T) {

	account1 := CreateRandomAccount(t)
	arg := ListEntriesParams{
		Limit:  3,
		Offset: 2,
	}
	for i := 0; i < 5; i++ {
		CreateRandomEntry(t, account1)
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
