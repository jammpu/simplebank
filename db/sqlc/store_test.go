package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// correr n transferencias concurrentes
	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {

			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: fromAccountID,
				ToAccountId:   toAccountID,
				Amount:        amount,
			})
			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	updateAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updateAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updateAccount1.Balance, updateAccount2.Balance)
	require.Equal(t, account1.Balance, updateAccount1.Balance)
	require.Equal(t, account2.Balance, updateAccount2.Balance)
}

//// checar resultados
//for i := 0; i < n; i++ {
//err := <-errs
//require.NoError(t, err)
//
//result := <-results
//require.NotEmpty(t, result)
//
////checar transferencia
//transfer := result.Transfer
//require.NotEmpty(t, transfer)
//require.Equal(t, account1.ID, transfer.FromAccountID)
//require.Equal(t, account2.ID, transfer.ToAccountID)
//require.Equal(t, amount, transfer.Amount)
//require.NotZero(t, transfer.ID)
//require.NotZero(t, transfer.CreatedAt)
//
////checar que los IDs de las transferencias existan
//_, err = store.GetTransfer(context.Background(), transfer.ID)
//require.NoError(t, err)
//
////checar entradas
//fromEntry := result.FromEntry
//require.NotEmpty(t, fromEntry)
//require.Equal(t, account1.ID, fromEntry.AccountID)
//require.Equal(t, -amount, fromEntry.Amount)
//require.NotZero(t, fromEntry.ID)
//require.NotZero(t, fromEntry.CreatedAt)
//
//_, err = store.GetEntry(context.Background(), fromEntry.ID)
//require.NoError(t, err)
//
//ToEntry := result.ToEntry
//require.NotEmpty(t, ToEntry)
//require.Equal(t, account2.ID, ToEntry.AccountID)
//require.Equal(t, amount, ToEntry.Amount)
//require.NotZero(t, ToEntry.ID)
//require.NotZero(t, ToEntry.CreatedAt)
//
//_, err = store.GetEntry(context.Background(), ToEntry.ID)
//require.NoError(t, err)
//
//// checar cuentas
//fromAccount := result.FromAccount
//require.NotEmpty(t, fromAccount)
//require.Equal(t, account1.ID, fromAccount.ID)
//
//toAccount := result.ToAccount
//require.NotEmpty(t, toAccount)
//require.Equal(t, account2.ID, toAccount.ID)
//
//// checar balance de cuenta
//fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
//diff1 := account1.Balance - fromAccount.Balance
//diff2 := toAccount.Balance - account2.Balance
//require.Equal(t, diff1, diff2)
//require.True(t, diff1 > 0)
//require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, 4 * amount.... n * amount
//
//k := int(diff1 / amount)
//require.True(t, k >= 1 && k <= n)
//require.NotContains(t, existed, k)
//existed[k] = true
