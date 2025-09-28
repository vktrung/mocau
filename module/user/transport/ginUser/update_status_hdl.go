package ginUser

import (
	"mocau-backend/common"
	"mocau-backend/module/user/biz"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ToggleUserStatus godoc
// @Summary Toggle user status
// @Description Toggle user status between active and inactive
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} common.Response "User status toggled successfully"
// @Failure 400 {object} common.Response "Invalid request"
// @Failure 401 {object} common.Response "Unauthorized"
// @Failure 404 {object} common.Response "User not found"
// @Router /users/{id}/toggle-status [put]
func ToggleUserStatus(business *biz.UpdateUserStatusBusiness) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy user ID từ URL parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// Gọi business logic để toggle status
		if err := business.ToggleUserStatus(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(gin.H{
			"message": "User status toggled successfully",
			"user_id": id,
		}))
	}
}
