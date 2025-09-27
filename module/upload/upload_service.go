package upload

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"mocau-backend/common"
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
	
	// Tạo thư mục media nếu chưa tồn tại
	uploadDir := "/media"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("không thể tạo thư mục media: %v", err)
	}
	
	// Tạo tên file mới với timestamp và extension
	newFileName := fmt.Sprintf("%d%s", time.Now().UTC().UnixNano(), ext)
	dst := fmt.Sprintf("%s/%s", uploadDir, newFileName)

	// Lưu file
	if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
		return nil, err
	}

	// Lấy base URL từ environment hoặc dùng default
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://160.250.5.71:3000"
	}

	// Tạo Image object với URL path
	imagePath := fmt.Sprintf("media/%s", newFileName)
	img := &common.Image{
		Id:        0,
		Url:       fmt.Sprintf("%s/%s", baseURL, imagePath),
		Width:     100,
		Height:    100,
		CloudName: "vps",
		Extension: ext,
	}

	return img, nil
}

// DeleteImage xóa file ảnh từ VPS filesystem
func DeleteImage(imageUrl string) error {
	if imageUrl == "" {
		return nil
	}

	// Lấy base URL từ environment hoặc dùng default
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://160.250.5.71:3000"
	}

	// Loại bỏ domain để lấy đường dẫn file
	urlPrefix := fmt.Sprintf("%s/media/", baseURL)
	if !strings.HasPrefix(imageUrl, urlPrefix) {
		return fmt.Errorf("invalid image URL format")
	}

	filename := strings.TrimPrefix(imageUrl, urlPrefix)
	filePath := fmt.Sprintf("/media/%s", filename)

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
