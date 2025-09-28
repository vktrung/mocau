package ginUser

import (
	"mocau-backend/common"
	"mocau-backend/module/user/biz"
	"mocau-backend/module/user/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update user profile information (full_name, phone, email)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param profile body model.UserUpdate true "Profile update request"
// @Success 200 {object} common.Response "User profile updated successfully"
// @Failure 400 {object} common.Response "Invalid request"
// @Failure 401 {object} common.Response "Unauthorized"
// @Failure 404 {object} common.Response "User not found"
// @Router /users/{id}/profile [put]
func UpdateProfile(business *biz.UpdateUserProfileBusiness) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy user ID từ URL parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// Parse request body
		var req model.UserUpdate
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// Gọi business logic
		if err := business.UpdateUserProfile(c.Request.Context(), id, &req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(gin.H{
			"message": "User profile updated successfully",
			"user_id": id,
		}))
	}
}
