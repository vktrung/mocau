package middleware

import (
	"context"
	"fmt"
	"mocau-backend/common"
	"mocau-backend/component/tokenprovider"
	"mocau-backend/module/user/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	// "Authorization": "Bearer {token}" or just "{token}"

	// If only one part, treat it as token directly
	if len(parts) == 1 {
		if strings.TrimSpace(parts[0]) == "" {
			return "", ErrWrongAuthHeader(nil)
		}
		return strings.TrimSpace(parts[0]), nil
	}

	// If multiple parts, expect "Bearer {token}" format
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return strings.TrimSpace(parts[1]), nil
}

func RequiredAuth(authStore AuthenStore, tokenProvider tokenprovider.Provider) func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := authStore.FindUser(
			c.Request.Context(),
			map[string]interface{}{"id": payload.UserId()},
		)
		if err != nil {
			panic(err)
		}

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
