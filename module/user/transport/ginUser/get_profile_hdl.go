package ginUser

import (
	"mocau-backend/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile godoc
// @Summary Get user profile
// @Description Get current user profile information
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} common.Response{data=model.User} "User profile retrieved successfully"
// @Failure 401 {object} common.Response "Unauthorized"
// @Router /profile [get]
func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessRes(u))
	}
}
