package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)


type Store interface {
	// TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error)
	SearchProductsTx(ctx context.Context, arg SearchParams) (SearchProductsResult, error)
	SearchFeedsTx(ctx context.Context, arg SearchParams) (SearchFeedsResult, error)
	Querier
} 

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool // in order to DB transaction
	*Queries	
} 

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries: New(connPool),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(context context.Context, callback func(*Queries) error) error {
	tx, err := store.connPool.Begin(context)
	if err != nil {
		return err
	}

	queries := New(tx)
	err = callback(queries)
	if err != nil {
		if rbErr := tx.Rollback(context); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err				
	}	
	return tx.Commit(context)
}	

type SearchParams struct {
	Query string `uri:"query" binding:"required"`
	Limit int32  `uri:"limit" binding:"required"`
	Offset int32 `uri:"offset" binding:"required"`
}
type SearchProductsResult struct {
	Products []GetProductsWithSellerByQueryRow `json:"products"`
}
func (store *SQLStore) SearchProductsTx(ctx context.Context, arg SearchParams) (SearchProductsResult, error) {
	var result SearchProductsResult		
	err := store.execTx(ctx, func(queries *Queries) error {		
		_, err := queries.UpsertSearchCount(ctx, arg.Query)
		if err != nil { return err }

		result.Products, err = queries.GetProductsWithSellerByQuery(ctx, GetProductsWithSellerByQueryParams{
			Column1: pgtype.Text{String: arg.Query, Valid: true},			
			Limit: arg.Limit,
			Offset: arg.Offset,			
		})
		if err != nil { return err }

		return nil
	})
	return result, err
}

type SearchFeedsResult struct {
	Feeds []GetFeedsWithAuthorByQueryRow `json:"feeds"`
}
func (store *SQLStore) SearchFeedsTx(ctx context.Context, arg SearchParams) (SearchFeedsResult, error) {
	var result SearchFeedsResult
	err := store.execTx(ctx, func (queries *Queries) error  {
		_, err := queries.UpsertSearchCount(ctx, arg.Query)
		if err != nil { return err }

		result.Feeds, err = queries.GetFeedsWithAuthorByQuery(ctx, GetFeedsWithAuthorByQueryParams{
			Column1: pgtype.Text{String: arg.Query, Valid: true},			
			Limit: arg.Limit,
			Offset: arg.Offset,		
		})
		if err != nil {return err}

		return nil
	})

	return result, err
}