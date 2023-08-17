package main

import (
	"database/sql"
	"log"
	"time"

	route "github.com/ChooseCruise/choosecruise-backend/api/route"
	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	"github.com/ChooseCruise/choosecruise-backend/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	conn, err := sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	store := repository.NewStore(conn)

	// timeout := time.Duration(env.ContextTimeout) * time.Second
	timeout := time.Second * 2

	gin := gin.Default()

	route.Setup(env, timeout, store, gin)

	gin.Run(env.ServerAddress)
}
