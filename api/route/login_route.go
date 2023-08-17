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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, store repository.Store, group *gin.RouterGroup, maker tokenutil.Maker) {
	ur := repository.NewUserRepository(store)
	sr := repository.NewSessionRepository(store)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, sr, timeout),
		Env:          env,
		Maker:        maker,
	}
	group.POST("/login", lc.Login)
}
