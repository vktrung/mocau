package ginBlog

import (
	"mocau-backend/common"
	"mocau-backend/module/blog/biz"
	"mocau-backend/module/blog/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListBlogs godoc
// @Summary List all blogs
// @Description Get list of all blogs
// @Tags blogs
// @Accept json
// @Produce json
// @Param status query string false "Filter by status (draft, published)"
// @Param author_id query int false "Filter by author ID"
// @Success 200 {object} common.Response{data=[]model.Blog}
// @Failure 500 {object} common.Response "Internal error"
// @Router /blogs [get]
func ListBlogs(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var filter = make(map[string]interface{})

		// Filter by status
		if status := c.Query("status"); status != "" {
			filter["status"] = status
		}

		// Filter by author_id
		if authorId := c.Query("author_id"); authorId != "" {
			filter["author_id"] = authorId
		}

		store := storage.NewSQLStore(db)
		b := biz.NewListBusiness(store)
		data, err := b.ListBlogs(c.Request.Context(), filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}
