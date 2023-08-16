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

	conn, err := sql.Open(env.DBHost, env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	db := repository.NewStore(conn)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
