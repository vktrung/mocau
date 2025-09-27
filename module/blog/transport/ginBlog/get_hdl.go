package ginBlog

import (
	"mocau-backend/common"
	"mocau-backend/module/blog/biz"
	"mocau-backend/module/blog/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetBlog godoc
// @Summary Get blog by ID
// @Description Get blog details by ID with author information
// @Tags blogs
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} common.Response{data=model.BlogWithAuthor}
// @Failure 404 {object} common.Response "Blog not found"
// @Router /blogs/{id} [get]
func GetBlog(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		b := biz.NewGetBusiness(store)
		data, err := b.GetBlogWithAuthor(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}
