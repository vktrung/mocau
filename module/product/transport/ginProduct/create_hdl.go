package ginProduct

import (
	"mocau-backend/common"
	"mocau-backend/module/product/biz"
	"mocau-backend/module/product/model"
	"mocau-backend/module/product/storage"
	"mocau-backend/module/upload"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with name, description, price, stock, category and image
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Product name"
// @Param description formData string false "Product description"
// @Param price formData number true "Product price"
// @Param stock formData integer true "Product stock"
// @Param category_id formData integer false "Category ID"
// @Param image formData file false "Product image"
// @Success 200 {object} common.Response{data=int} "Product created successfully"
// @Failure 400 {object} common.Response "Invalid request data"
// @Router /products [post]
func CreateProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ProductCreate

		// Bind form data manually for multipart/form-data
		data.Name = c.PostForm("name")
		data.Description = c.PostForm("description")
		
		// Parse price
		if priceStr := c.PostForm("price"); priceStr != "" {
			if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
				data.Price = price
			}
		}
		
		// Parse stock
		if stockStr := c.PostForm("stock"); stockStr != "" {
			if stock, err := strconv.Atoi(stockStr); err == nil {
				data.Stock = stock
			}
		}
		
		// Parse category_id
		if categoryIdStr := c.PostForm("category_id"); categoryIdStr != "" {
			if categoryId, err := strconv.Atoi(categoryIdStr); err == nil {
				data.CategoryId = &categoryId
			}
		}

		// Xử lý upload ảnh nếu có
		if img, err := upload.UploadImage(c, "image"); err != nil {
			panic(common.ErrInvalidRequest(err))
		} else if img != nil {
			data.Image = img
		}

		store := storage.NewSQLStore(db)
		b := biz.NewCreateBusiness(store)
		if err := b.CreateProduct(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
