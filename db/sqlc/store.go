package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	connPool *pgxpool.Pool // in order to DB transaction
} 

func NewStore(connPool *pgxpool.Pool) *Store {
	return &Store{
		connPool:      connPool,
		Queries: New(connPool),
	}
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer 	Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount  	Account  `json:"to_account"`
	FromEntry  	Entry    `json:"from_entry"`
	ToEntry		Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one user to another
// It creates a transfer record and update accounts balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error) {
	var result TransferTxResult	
	var err error
	store.execTx(ctx, func(queries *Queries) error {														
		result.Transfer, err = queries.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})		
		if err != nil { return err }		
		
		result.FromEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})		
		if err != nil { return err }		
		result.ToEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})		
		if err != nil { return err }

		// OLD WAY
		// fromAccount, getAccountErr := queries.GetAccountForUpdate(ctx, arg.FromAccountID)		
		// if getAccountErr != nil { return getAccountErr }
		// toAccount, getAccountErr := queries.GetAccountForUpdate(ctx, arg.ToAccountID)							
		// if getAccountErr != nil { return getAccountErr }
		
		// result.FromAccount, err = queries.UpdateAccount(ctx, UpdateAccountParams{
		// 	ID:      arg.FromAccountID,
		// 	Owner:  fromAccount.Owner,
		// 	Currency: fromAccount.Currency,
		// 	Balance: fromAccount.Balance - arg.Amount,			
		// })
		// if err != nil {	return err }
		
		// result.ToAccount, err = queries.UpdateAccount(ctx, UpdateAccountParams{
		// 	ID:      arg.ToAccountID,
		// 	Owner: toAccount.Owner,
		// 	Currency: toAccount.Currency,
		// 	Balance: toAccount.Balance + arg.Amount,
		// })

		// result.FromAccount, err = queries.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		// 	ID:      arg.FromAccountID,
		// 	Amount: -arg.Amount,			
		// })
		// if err != nil {	return err }
		
		// result.ToAccount, err = queries.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		// 	ID:      arg.ToAccountID,
		// 	Amount: arg.Amount,
		// })

		if arg.FromAccountID < arg.ToAccountID {
 			result.FromAccount, result.ToAccount, err = sendMonny(ctx, queries, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = sendMonny(ctx, queries, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}
		if err != nil {	return err }

		return nil
	})	
	return result, err
}

// in order to DB transaction with no deadlock
func sendMonny(
	ctx context.Context,
	queries *Queries,
	account1 int64,	
	amount1 int64,	
	account2 int64,
	amount2 int64,
) (fromAccount Account, toAccount Account, err error) {
	fromAccount, err = queries.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:	  account1,
		Amount: amount1,
	})
	if err != nil {	return fromAccount, toAccount, err }
	toAccount, err = queries.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:	  account2,
		Amount: amount2,
	})	
	if err != nil {	return fromAccount, toAccount, err }
	return fromAccount, toAccount, nil
}


type CreatePostTxParams struct {
	UserID   int64 			`json:"user_id"`
	Title    string 		`json:"title"`
	Content  string 		`json:"content"`
	Media    []string 		`json:"media"`
	Location string 		`json:"location"`
	ProductID pgtype.Int8   `json:"product_id"` // nullable ProductID field 
}

type CreatePostTxResult struct {
	Post    Post
	User    User
	Product Product
}

func (store *Store) CreatePostTx(ctx context.Context, arg CreatePostTxParams) (CreatePostTxResult, error) {
	var result CreatePostTxResult	
	var err error
	
	err = store.execTx(ctx, func(queries *Queries) error {
		result.User, err = queries.GetUser(ctx, arg.UserID)
		if arg.ProductID.Valid { 
			result.Product, err = queries.GetProduct(ctx, arg.ProductID.Int64)			
			if err != nil {
				return err
			}
		} 
		result.Post, err = queries.CreatePost(ctx, CreatePostParams{
			UserID:    result.User.ID,			
			Title:     arg.Title,
			Content:   arg.Content,
			Media:     arg.Media,
			Location:  pgtype.Text{String: arg.Location, Valid: true},		
			ProductID: arg.ProductID, // nullable ProductID field	
		})
		if err != nil {
			return err			
		}
		return nil
	})	

	return result, err
}