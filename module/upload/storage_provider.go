package upload

import (
	"github.com/gin-gonic/gin"
	"mocau-backend/common"
)

// StorageProvider interface for different storage backends
type StorageProvider interface {
	UploadImage(c *gin.Context, fieldName string) (*common.Image, error)
	DeleteImage(imageUrl string) error
	GetProviderName() string
}
