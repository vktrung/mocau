package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mocau-backend/common"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// UploadImage xử lý upload ảnh và trả về Image object
func UploadImage(c *gin.Context, fieldName string) (*common.Image, error) {
	fileHeader, err := c.FormFile(fieldName)
	if err != nil {
		return nil, err
	}

	// Lấy extension từ file gốc
	ext := filepath.Ext(fileHeader.Filename)
	ext = strings.ToLower(ext) // Chuyển về lowercase để so sánh
	
	// Kiểm tra định dạng file chỉ cho phép PNG và JPG
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		return nil, fmt.Errorf("chỉ cho phép file PNG hoặc JPG")
	}
	
	// Tạo tên file mới với timestamp và extension
	newFileName := fmt.Sprintf("%d%s", time.Now().UTC().UnixNano(), ext)
	dst := fmt.Sprintf("static/%s", newFileName)

	// Lưu file
	if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
		return nil, err
	}

	// Tạo Image object
	img := &common.Image{
		Id:        0,
		Url:       dst,
		Width:     100,
		Height:    100,
		CloudName: "local",
		Extension: "",
	}

	// Tạo URL đầy đủ với domain VPS
	img.Fulfill("http://160.250.5.71:3000")

	return img, nil
}

// DeleteImage xóa file ảnh từ filesystem
func DeleteImage(imageUrl string) error {
	if imageUrl == "" {
		return nil
	}

	// Loại bỏ domain để lấy đường dẫn file
	filePath := imageUrl
	if len(imageUrl) > len("http://160.250.5.71:3000/") {
		filePath = imageUrl[len("http://160.250.5.71:3000/"):]
	}

	// Kiểm tra file có tồn tại không
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil // File không tồn tại, không cần xóa
	}

	// Xóa file
	return os.Remove(filePath)
}

// DeleteImageFromProduct xóa ảnh từ product object
func DeleteImageFromProduct(product *common.Image) error {
	if product == nil {
		return nil
	}
	return DeleteImage(product.Url)
}
