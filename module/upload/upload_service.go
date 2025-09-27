package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mocau-backend/common"
	"os"
	"time"
)

// UploadImage xử lý upload ảnh và trả về Image object
func UploadImage(c *gin.Context, fieldName string) (*common.Image, error) {
	fileHeader, err := c.FormFile(fieldName)
	if err != nil {
		return nil, err
	}

	// Tạo tên file unique
	dst := fmt.Sprintf("static/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)

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

	// Tạo URL đầy đủ
	img.Fulfill("http://localhost:3000")

	return img, nil
}

// DeleteImage xóa file ảnh từ filesystem
func DeleteImage(imageUrl string) error {
	if imageUrl == "" {
		return nil
	}

	// Loại bỏ domain để lấy đường dẫn file
	filePath := imageUrl
	if len(imageUrl) > len("http://localhost:3000/") {
		filePath = imageUrl[len("http://localhost:3000/"):]
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
