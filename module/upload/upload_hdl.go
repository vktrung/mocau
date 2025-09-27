package upload

import (
	"mocau-backend/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Upload(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// Upload image directly to VPS
		img, err := UploadImage(c, "file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(img))
	}
}
