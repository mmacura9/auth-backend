package route

import (
	"time"

	"github.com/ChooseCruise/choosecruise-backend/api/controller"
	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/ChooseCruise/choosecruise-backend/repository"
	"github.com/ChooseCruise/choosecruise-backend/usecase"
	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db repository.Store, group *gin.RouterGroup, maker tokenutil.Maker) {
	ur := repository.NewUserRepository(db)
	sr := repository.NewSessionRepository(db)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, sr, timeout),
		Env:                 env,
		Maker:               maker,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
