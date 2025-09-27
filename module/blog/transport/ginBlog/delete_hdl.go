package ginBlog

import (
	"mocau-backend/common"
	"mocau-backend/module/blog/biz"
	"mocau-backend/module/blog/storage"
	userModel "mocau-backend/module/user/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteBlog godoc
// @Summary Soft-delete a blog by id
// @Description Set deleted_at timestamp instead of removing the record. Only blog author can delete their own blog.
// @Tags blogs
// @Produce json
// @Security BearerAuth
// @Param id path int true "Blog ID"
// @Success 200 {object} common.Response{data=bool}
// @Failure 400 {object} common.AppError
// @Failure 403 {object} common.Response "Forbidden - not blog author"
// @Failure 404 {object} common.AppError
// @Router /blogs/{id} [delete]
func DeleteBlog(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Lấy thông tin user từ context (đã được set bởi middleware auth)
		user := c.MustGet(common.CurrentUser)
		userObj, ok := user.(*userModel.User)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}

		// Kiểm tra quyền sở hữu blog
		store := storage.NewSQLStore(db)
		currentBlog, err := store.GetBlog(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			if err == common.RecordNotFound {
				panic(common.ErrEntityNotFound("Blog", err))
			}
			panic(common.ErrCannotGetEntity("Blog", err))
		}

		// Chỉ cho phép tác giả xóa blog của mình
		if currentBlog.AuthorId != userObj.Id {
			panic(common.ErrNoPermission(nil))
		}

		b := biz.NewDeleteBusiness(store)
		if err := b.DeleteBlog(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
