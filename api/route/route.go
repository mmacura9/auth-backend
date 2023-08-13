package route

import (
	"log"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/api/middleware"
	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/ChooseCruise/choosecruise-backend/repository"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db repository.Store, gin *gin.Engine) {
	tokenMaker, err := tokenutil.NewPastoMaker(env.RefreshTokenSecret)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter, tokenMaker)
	NewLoginRouter(env, timeout, db, publicRouter, tokenMaker)
	NewRefreshTokenRouter(env, timeout, db, publicRouter, tokenMaker)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.AuthMiddleware(tokenMaker))
	// All Private APIs

}
