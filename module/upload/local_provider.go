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

type LocalProvider struct{}

func NewLocalProvider() *LocalProvider {
	return &LocalProvider{}
}

func (lp *LocalProvider) UploadImage(c *gin.Context, fieldName string) (*common.Image, error) {
	fileHeader, err := c.FormFile(fieldName)
	if err != nil {
		return nil, err
	}

	// Validate file extension
	ext := filepath.Ext(fileHeader.Filename)
	ext = strings.ToLower(ext)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		return nil, fmt.Errorf("chỉ cho phép file PNG hoặc JPG")
	}

	// Create uploads directory
	uploadDir := "/media"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("không thể tạo thư mục media: %v", err)
	}

	// Generate unique filename
	newFileName := fmt.Sprintf("%d%s", time.Now().UTC().UnixNano(), ext)
	dst := fmt.Sprintf("%s/%s", uploadDir, newFileName)

	// Save file
	if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
		return nil, err
	}

	// Get base URL
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://160.250.5.71:3000"
	}

	// Create Image object
	imagePath := fmt.Sprintf("media/%s", newFileName)
	img := &common.Image{
		Id:        0,
		Url:       fmt.Sprintf("%s/%s", baseURL, imagePath),
		Width:     100,
		Height:    100,
		CloudName: "local",
		Extension: ext,
	}

	return img, nil
}

func (lp *LocalProvider) DeleteImage(imageUrl string) error {
	if imageUrl == "" {
		return nil
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://160.250.5.71:3000"
	}

	urlPrefix := fmt.Sprintf("%s/media/", baseURL)
	if !strings.HasPrefix(imageUrl, urlPrefix) {
		return fmt.Errorf("invalid image URL format")
	}

	filename := strings.TrimPrefix(imageUrl, urlPrefix)
	filePath := fmt.Sprintf("/media/%s", filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(filePath)
}

func (lp *LocalProvider) GetProviderName() string {
	return "local"
}
