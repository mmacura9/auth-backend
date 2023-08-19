package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	env := bootstrap.NewEnv()
	conn, err := sql.Open(env.DBDriver, env.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
