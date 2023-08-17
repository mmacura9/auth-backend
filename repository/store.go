package repository

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
)

type Store interface {
	db.Querier
	execTx(ctx context.Context, fn func(*db.Queries) error) error
}

type SQLStore struct {
	*db.Queries
	sqldb *sql.DB
}

func NewStore(sqldb *sql.DB) Store {
	return &SQLStore{Queries: db.New(sqldb), sqldb: sqldb}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := store.sqldb.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx err %v, rollback error %v", err, rberr)
		}
		return err
	}

	return tx.Commit()
}
