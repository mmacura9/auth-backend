package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey  = "authorization"
	authTypeBearer = "bearer"
	authPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker tokenutil.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authHeaderKey)
		if len(authHeader) == 0 {
			err := errors.New("Authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewErrorResponse(err.Error()))
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("Authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewErrorResponse(err.Error()))
			return
		}

		authType := strings.ToLower(fields[0])

		if authType != authTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewErrorResponse(err.Error()))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			err := errors.New("Not verified")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewErrorResponse(err.Error()))
			return
		}
		ctx.Set(authPayloadKey, payload)
		ctx.Next()
	}
}
