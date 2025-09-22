package upload

import (
	"mocau-backend/common"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Upload(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dst := fmt.Sprintf("static/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)

		if err := c.SaveUploadedFile(fileHeader, dst); err != nil {

		}

		img := &common.Image{
			Id:        0,
			Url:       dst,
			Width:     100,
			Height:    100,
			CloudName: "local",
			Extension: "",
		}

		img.Fulfill("http://localhost:3000")

		c.JSON(http.StatusOK, common.SimpleSuccessRes(img))
	}
}
