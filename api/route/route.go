package route

import (
	"log"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/api/middleware"
	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, store db.Store, gin *gin.Engine) {
	tokenMaker, err := tokenutil.NewPasetoMaker(env.RefreshTokenSecret)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}
	apiRouter := gin.Group("api")
	v1Router := apiRouter.Group("v1")
	// All Public APIs
	NewSignupRouter(env, timeout, store, v1Router, tokenMaker)

	authV1Router := v1Router.Group("auth")
	// Middleware to verify AccessToken
	authV1Router.Use(middleware.AuthMiddleware(tokenMaker))
	// All Private APIs

}
