package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"

	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
)

var testStore Store

func TestMain(m *testing.M) {
	env := bootstrap.LoadEnv("../..")

	connPool, err := pgxpool.New(context.Background(), env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
