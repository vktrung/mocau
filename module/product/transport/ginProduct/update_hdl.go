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

// UpdateProduct godoc
// @Summary Update a product by ID
// @Description Update product fields by ID including image
// @Tags products
// @Accept multipart/form-data
// @Produce json
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
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            panic(common.ErrInvalidRequest(err))
        }

        var data model.ProductUpdate
        
        // Bind form data
        if err := c.ShouldBind(&data); err != nil {
            panic(common.ErrInvalidRequest(err))
        }

        // Xử lý upload ảnh nếu có
        if img, err := upload.UploadImage(c, "image"); err == nil {
            data.Image = img
        }

        store := storage.NewSQLStore(db)
        b := biz.NewUpdateBusiness(store)
        if err := b.UpdateProduct(c.Request.Context(), id, &data); err != nil {
            panic(err)
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
    }
}


