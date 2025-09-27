package upload

import (
	"mocau-backend/common"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
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

		// Lấy extension từ file gốc
		ext := filepath.Ext(fileHeader.Filename)
		ext = strings.ToLower(ext) // Chuyển về lowercase để so sánh
		
		// Kiểm tra định dạng file chỉ cho phép PNG và JPG
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			panic(common.ErrInvalidRequest(fmt.Errorf("chỉ cho phép file PNG hoặc JPG")))
		}
		
		// Tạo tên file mới với timestamp và extension
		newFileName := fmt.Sprintf("%d%s", time.Now().UTC().UnixNano(), ext)
		dst := fmt.Sprintf("static/%s", newFileName)

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

		img.Fulfill("http://160.250.5.71:3000")

		c.JSON(http.StatusOK, common.SimpleSuccessRes(img))
	}
}
