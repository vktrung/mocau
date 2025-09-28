package ginUser

import (
	"mocau-backend/common"
	"mocau-backend/module/user/biz"
	"mocau-backend/module/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListUsers godoc
// @Summary List all users
// @Description Get a list of all users with optional filtering
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param status query string false "Filter by status"
// @Param role query string false "Filter by role"
// @Success 200 {object} common.Response{data=[]model.User} "Users retrieved successfully"
// @Failure 400 {object} common.Response "Bad request"
// @Failure 401 {object} common.Response "Unauthorized"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /users [get]
func ListUsers(biz *biz.ListUserBusiness) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.UserFilter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		result, err := biz.ListUsers(c.Request.Context(), &filter)

		if err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(result))
	}
}
