package upload

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"mocau-backend/common"
)

type CloudinaryProvider struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryProvider() (*CloudinaryProvider, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return nil, fmt.Errorf("missing Cloudinary credentials in environment variables")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}

	return &CloudinaryProvider{cld: cld}, nil
}

func (cp *CloudinaryProvider) UploadImage(c *gin.Context, fieldName string) (*common.Image, error) {
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

	// Open uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer file.Close()

	// Generate unique public ID
	publicID := fmt.Sprintf("mocau/%d_%s", time.Now().UTC().UnixNano(), strings.TrimSuffix(fileHeader.Filename, ext))
	
	// Upload to Cloudinary
	result, err := cp.cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:     publicID,
		Folder:       "mocau",
		ResourceType: "image",
		Format:       strings.TrimPrefix(ext, "."),
		Transformation: "q_auto,f_auto",
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to upload to Cloudinary: %v", err)
	}

	// Create Image object
	img := &common.Image{
		Id:        0,
		Url:       result.SecureURL,
		Width:     result.Width,
		Height:    result.Height,
		CloudName: "cloudinary",
		Extension: ext,
	}

	return img, nil
}

func (cp *CloudinaryProvider) DeleteImage(imageUrl string) error {
	if imageUrl == "" {
		return nil
	}

	// Extract public ID from Cloudinary URL
	publicID := cp.extractPublicIDFromURL(imageUrl)
	if publicID == "" {
		return fmt.Errorf("invalid Cloudinary URL")
	}

	_, err := cp.cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "image",
	})

	return err
}

func (cp *CloudinaryProvider) GetProviderName() string {
	return "cloudinary"
}

func (cp *CloudinaryProvider) extractPublicIDFromURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) < 7 {
		return ""
	}

	for i, part := range parts {
		if part == "upload" {
			if i+1 < len(parts) {
				publicIDWithExt := parts[len(parts)-1]
				if dotIndex := strings.LastIndex(publicIDWithExt, "."); dotIndex != -1 {
					return "mocau/" + publicIDWithExt[:dotIndex]
				}
				return "mocau/" + publicIDWithExt
			}
		}
	}

	return ""
}
