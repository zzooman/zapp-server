package db

// func createRandomAccount(t *testing.T) Account {
// 	user := createRandomUser(t)

// 	arg := CreateAccountParams{
// 		Owner:    user.Username,
// 		Balance:  int64(1000),
// 		Currency: "USD",
// 	}

// 	account, err := testStore.CreateAccount(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account)

// 	require.Equal(t, arg.Owner, account.Owner)
// 	require.Equal(t, arg.Balance, account.Balance)
// 	require.Equal(t, arg.Currency, account.Currency)

// 	return account
// }

// func TestStore_TransferTx(t *testing.T) {
// 	fromAccount := createRandomAccount(t)
// 	toAccount := createRandomAccount(t)

// 	n := 5
// 	amount := int64(100)
// 	errs := make(chan error)
// 	results := make(chan TransferTxResult)

// 	for i := 0; i < n; i++ {
// 		go func() {
// 			result, err := testStore.TransferTx(context.Background(), CreateTransferParams{
// 				FromAccountID: fromAccount.ID,
// 				ToAccountID:  toAccount.ID,
// 				Amount: amount,
// 			})
// 			errs <- err
// 			results <- result
// 		}()
// 	}

// 	existed := make(map[int]bool)
// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)

// 		result := <-results
// 		require.NotEmpty(t, result)

// 		transfer := result.Transfer
// 		require.NotEmpty(t, transfer)
// 		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
// 		require.Equal(t, toAccount.ID, transfer.ToAccountID)
// 		require.Equal(t, amount, transfer.Amount)

// 		_, err = testStore.GetAccount(context.Background(), fromAccount.ID)
// 		require.NoError(t, err)

// 		fromEntry := result.FromEntry
// 		require.NotEmpty(t, fromEntry)
// 		require.Equal(t, fromAccount.ID, fromEntry.AccountID)
// 		require.Equal(t, -amount, fromEntry.Amount)

// 		_, err = testStore.GetEntry(context.Background(), fromEntry.ID)
// 		require.NoError(t, err)

// 		toEntry := result.ToEntry
// 		require.NotEmpty(t, toEntry)
// 		require.Equal(t, toAccount.ID, toEntry.AccountID)
// 		require.Equal(t, amount, toEntry.Amount)

// 		_, err = testStore.GetEntry(context.Background(), toEntry.ID)
// 		require.NoError(t, err)

// 		// check account
// 		require.NotEmpty(t, result.FromAccount)
// 		require.Equal(t, fromAccount.ID, result.FromAccount.ID)
// 		require.NotEmpty(t, result.ToAccount)
// 		require.Equal(t, toAccount.ID, result.ToAccount.ID)

// 		// check account balance
// 		diff1 := fromAccount.Balance - result.FromAccount.Balance
// 		diff2 := result.ToAccount.Balance - toAccount.Balance
// 		require.Equal(t, diff1, diff2)
// 		require.True(t, diff1 > 0)
// 		require.True(t, diff1%amount == 0)

// 		k := int(diff1 / amount)
// 		require.True(t, k >= 1 && k <= n)
// 		require.NotContains(t, existed, k)
// 		existed[k] = true
// 	}

// 	updatedFromAccount, err := testStore.GetAccount(context.Background(), fromAccount.ID)
// 	require.NoError(t, err)
// 	updatedToAccount, err := testStore.GetAccount(context.Background(), toAccount.ID)
// 	require.NoError(t, err)

// 	require.Equal(t, fromAccount.Balance-int64(int64(n)*amount), updatedFromAccount.Balance)
// 	require.Equal(t, toAccount.Balance+int64(int64(n)*amount), updatedToAccount.Balance)
// }

// func TestStore_TransferTxDeadlock(t *testing.T) {
// 	fromAccount := createRandomAccount(t)
// 	toAccount := createRandomAccount(t)

// 	n := 10
// 	amount := int64(100)
// 	errs := make(chan error)

// 	for i := 0; i < n; i++ {
// 		go func() {
// 			fromAccountId := fromAccount.ID
// 			toAccountId := toAccount.ID
// 			if i%2 == 1 {
// 				fromAccountId, toAccountId = toAccountId, fromAccountId
// 			}
// 			_, err := testStore.TransferTx(context.Background(), CreateTransferParams{
// 				FromAccountID: fromAccountId,
// 				ToAccountID:  toAccountId,
// 				Amount: amount,
// 			})
// 			errs <- err
// 		}()
// 	}

// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)
// 	}

// 	updatedFromAccount, err := testStore.GetAccount(context.Background(), fromAccount.ID)
// 	require.NoError(t, err)
// 	updatedToAccount, err := testStore.GetAccount(context.Background(), toAccount.ID)
// 	require.NoError(t, err)

// 	require.Equal(t, fromAccount.Balance, updatedFromAccount.Balance)
// 	require.Equal(t, toAccount.Balance, updatedToAccount.Balance)
// }

