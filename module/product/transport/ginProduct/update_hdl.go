package ginProduct

import (
    "errors"
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

// UpdateProduct godoc
// @Summary Update a product by ID
// @Description Update product fields by ID including image
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param name formData string false "Product name"
// @Param description formData string false "Product description"
// @Param price formData number false "Product price"
// @Param stock formData integer false "Product stock"
// @Param category_id formData integer false "Category ID"
// @Param image formData file false "Product image"
// @Success 200 {object} common.Response{data=bool}
// @Failure 404 {object} common.Response "Product not found"
// @Router /products/{id} [put]
func UpdateProduct(db *gorm.DB) func(*gin.Context) {
    return func(c *gin.Context) {
        // Validate ID sớm
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil || id <= 0 {
            panic(common.ErrInvalidRequest(err))
        }

        var data model.ProductUpdate
        
        // Bind form data manually for multipart/form-data với validation
        if name := c.PostForm("name"); name != "" {
            if len(name) > 255 {
                panic(common.ErrInvalidRequest(errors.New("name too long")))
            }
            data.Name = &name
        }
        
        if description := c.PostForm("description"); description != "" {
            data.Description = &description
        }
        
        // Parse price với validation
        if priceStr := c.PostForm("price"); priceStr != "" {
            if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
                if price < 0 {
                    panic(common.ErrInvalidRequest(errors.New("price cannot be negative")))
                }
                data.Price = &price
            } else {
                panic(common.ErrInvalidRequest(errors.New("invalid price format")))
            }
        }
        
        // Parse stock với validation
        if stockStr := c.PostForm("stock"); stockStr != "" {
            if stock, err := strconv.Atoi(stockStr); err == nil {
                if stock < 0 {
                    panic(common.ErrInvalidRequest(errors.New("stock cannot be negative")))
                }
                data.Stock = &stock
            } else {
                panic(common.ErrInvalidRequest(errors.New("invalid stock format")))
            }
        }
        
        // Parse category_id với validation
        if categoryIdStr := c.PostForm("category_id"); categoryIdStr != "" {
            if categoryId, err := strconv.Atoi(categoryIdStr); err == nil {
                if categoryId <= 0 {
                    panic(common.ErrInvalidRequest(errors.New("invalid category_id")))
                }
                data.CategoryId = &categoryId
            } else {
                panic(common.ErrInvalidRequest(errors.New("invalid category_id format")))
            }
        }

        // Xử lý upload ảnh nếu có (chỉ khi cần thiết)
        if img, err := upload.UploadImage(c, "image"); err != nil {
            panic(common.ErrInvalidRequest(err))
        } else if img != nil {
            data.Image = img
        }

        // Kiểm tra xem có ít nhất một field được update không
        if data.Name == nil && data.Description == nil && data.Price == nil && 
           data.Stock == nil && data.CategoryId == nil && data.Image == nil {
            panic(common.ErrInvalidRequest(errors.New("no fields to update")))
        }

        store := storage.NewSQLStore(db)
        b := biz.NewUpdateBusiness(store)
        if err := b.UpdateProduct(c.Request.Context(), id, &data); err != nil {
            panic(err)
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
    }
}


