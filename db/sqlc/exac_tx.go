package db

import (
	"context"
	"fmt"
)

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