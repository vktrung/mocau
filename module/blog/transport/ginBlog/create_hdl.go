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

// CreateBlog godoc
// @Summary Create a new blog
// @Description Create a new blog with title, content, image and status. Author will be set automatically from logged in user.
// @Tags blogs
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Blog title"
// @Param content formData string true "Blog HTML content"
// @Param status formData string false "Blog status (draft or published)"
// @Param image formData file false "Blog image"
// @Success 200 {object} common.Response{data=bool} "Blog created successfully"
// @Failure 400 {object} common.Response "Invalid request data"
// @Router /blogs [post]
func CreateBlog(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.BlogCreate

		// Lấy thông tin user từ context (đã được set bởi middleware auth)
		user := c.MustGet(common.CurrentUser)
		userObj, ok := user.(*userModel.User)
		if !ok {
			panic(common.ErrUnauthorized(nil))
		}

		// Bind form data manually for multipart/form-data
		data.Title = c.PostForm("title")
		data.Content = c.PostForm("content")
		
		// Tự động set author_id từ user đang đăng nhập
		data.AuthorId = userObj.Id
		
		// Parse status (default to draft if empty)
		data.Status = c.PostForm("status")
		if data.Status == "" {
			data.Status = "draft"
		}

		// Xử lý upload ảnh nếu có
		if img, err := upload.UploadImage(c, "image"); err != nil {
			panic(common.ErrInvalidRequest(err))
		} else if img != nil {
			data.Image = img
		}

		store := storage.NewSQLStore(db)
		b := biz.NewCreateBusiness(store)
		if err := b.CreateBlog(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
