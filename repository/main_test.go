package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"

	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/ChooseCruise/choosecruise-backend/domain"
	_ "github.com/lib/pq"
)

var testStore db.Store
var userRep domain.UserRepository
var sessionRep domain.SessionRepository

func TestMain(m *testing.M) {
	env := bootstrap.LoadEnv("..")

	connPool, err := pgxpool.New(context.Background(), env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testStore = db.NewStore(connPool)
	userRep = NewUserRepository(testStore)
	sessionRep = NewSessionRepository(testStore)

	os.Exit(m.Run())
}
