package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	execTx(ctx context.Context, fn func(*Queries) error) error
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{Queries: New(db), db: db}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx err %v, rollback error %v", err, rberr)
		}
		return err
	}

	return tx.Commit()
}
