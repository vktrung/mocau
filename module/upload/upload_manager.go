package upload

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"mocau-backend/common"
)

type UploadManager struct {
	provider StorageProvider
}

// NewUploadManager creates upload manager based on STORAGE_TYPE environment variable
func NewUploadManager() (*UploadManager, error) {
	storageType := strings.ToLower(os.Getenv("STORAGE_TYPE"))
	
	var provider StorageProvider
	var err error

	switch storageType {
	case "cloudinary":
		provider, err = NewCloudinaryProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize Cloudinary: %v", err)
		}
		
	case "local", "":
		// Default to local storage
		provider = NewLocalProvider()
		
	default:
		return nil, fmt.Errorf("unsupported storage type: %s. Supported: cloudinary, local", storageType)
	}

	return &UploadManager{provider: provider}, nil
}

// UploadImage uploads image using configured provider
func (um *UploadManager) UploadImage(c *gin.Context, fieldName string) (*common.Image, error) {
	return um.provider.UploadImage(c, fieldName)
}

// DeleteImage deletes image using configured provider
func (um *UploadManager) DeleteImage(imageUrl string) error {
	return um.provider.DeleteImage(imageUrl)
}

// GetProviderName returns the name of current storage provider
func (um *UploadManager) GetProviderName() string {
	return um.provider.GetProviderName()
}
