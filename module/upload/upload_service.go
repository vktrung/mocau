package upload

import (
	"github.com/gin-gonic/gin"
	"mocau-backend/common"
)

// UploadImage xử lý upload ảnh và trả về Image object
func UploadImage(c *gin.Context, fieldName string) (*common.Image, error) {
	manager, err := NewUploadManager()
	if err != nil {
		return nil, err
	}
	
	return manager.UploadImage(c, fieldName)
}

// DeleteImage xóa file ảnh
func DeleteImage(imageUrl string) error {
	manager, err := NewUploadManager()
	if err != nil {
		return err
	}
	
	return manager.DeleteImage(imageUrl)
}

// DeleteImageFromProduct xóa ảnh từ product object
func DeleteImageFromProduct(product *common.Image) error {
	if product == nil {
		return nil
	}
	return DeleteImage(product.Url)
}
