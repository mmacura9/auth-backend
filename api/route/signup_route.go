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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db repository.Store, group *gin.RouterGroup, maker tokenutil.Maker) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
		Maker:         maker,
	}
	group.POST("/signup", sc.Signup)
}