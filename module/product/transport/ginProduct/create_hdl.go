package ginProduct

import (
    "mocau-backend/common"
    "mocau-backend/module/product/biz"
    "mocau-backend/module/product/model"
    "mocau-backend/module/product/storage"
    "mocau-backend/module/upload"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with name, description, price, stock, category and image
// @Tags products
// @Accept multipart/form-data
// @Produce json
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
        
        // Bind form data
        if err := c.ShouldBind(&data); err != nil {
            panic(common.ErrInvalidRequest(err))
        }

        // Xử lý upload ảnh nếu có
        if img, err := upload.UploadImage(c, "image"); err == nil {
            data.Image = img
        }

        store := storage.NewSQLStore(db)
        b := biz.NewCreateBusiness(store)
        if err := b.CreateProduct(c.Request.Context(), &data); err != nil {
            panic(err)
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(data.Id))
    }
}


