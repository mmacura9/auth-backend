package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/api/route"
	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	connPool, err := pgxpool.New(context.Background(), env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(connPool)

	timeout := time.Second * 2

	gin := gin.Default()

	route.Setup(env, timeout, store, gin)

	gin.Run(env.ServerAddress)
}
