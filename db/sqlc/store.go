package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	sqldb *sql.DB
}

func NewStore(sqldb *sql.DB) Store {
	return &SQLStore{Queries: New(sqldb), sqldb: sqldb}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.sqldb.BeginTx(ctx, nil)
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
