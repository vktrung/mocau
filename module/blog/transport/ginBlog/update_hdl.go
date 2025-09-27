package ginBlog

import (
	"mocau-backend/common"
	"mocau-backend/module/blog/biz"
	"mocau-backend/module/blog/model"
	"mocau-backend/module/blog/storage"
	"mocau-backend/module/upload"
	userModel "mocau-backend/module/user/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateBlog godoc
// @Summary Update a blog by ID
// @Description Update blog fields by ID including image. Only blog author can update their own blog.
// @Tags blogs
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Blog ID"
// @Param title formData string false "Blog title"
// @Param content formData string false "Blog HTML content"
// @Param status formData string false "Blog status (draft or published)"
// @Param image formData file false "Blog image"
// @Success 200 {object} common.Response{data=bool}
// @Failure 403 {object} common.Response "Forbidden - not blog author"
// @Failure 404 {object} common.Response "Blog not found"
// @Router /blogs/{id} [put]
func UpdateBlog(db *gorm.DB) func(*gin.Context) {
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

		// Chỉ cho phép tác giả chỉnh sửa blog của mình
		if currentBlog.AuthorId != userObj.Id {
			panic(common.ErrNoPermission(nil))
		}

		var data model.BlogUpdate
		
		// Bind form data manually for multipart/form-data
		if title := c.PostForm("title"); title != "" {
			data.Title = &title
		}
		if content := c.PostForm("content"); content != "" {
			data.Content = &content
		}
		if status := c.PostForm("status"); status != "" {
			data.Status = &status
		}

		// Xử lý upload ảnh nếu có
		if img, err := upload.UploadImage(c, "image"); err != nil {
			panic(common.ErrInvalidRequest(err))
		} else if img != nil {
			data.Image = img
		}

		b := biz.NewUpdateBusiness(store)
		if err := b.UpdateBlog(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
