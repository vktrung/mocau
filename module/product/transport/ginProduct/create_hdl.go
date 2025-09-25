package ginProduct

import (
    "mocau-backend/common"
    "mocau-backend/module/product/biz"
    "mocau-backend/module/product/model"
    "mocau-backend/module/product/storage"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with name, description, price, stock, category
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.ProductCreate true "Product data"
// @Success 200 {object} common.Response{data=int} "Product created successfully"
// @Failure 400 {object} common.Response "Invalid request data"
// @Router /products [post]
func CreateProduct(db *gorm.DB) func(*gin.Context) {
    return func(c *gin.Context) {
        var data model.ProductCreate
        if err := c.ShouldBind(&data); err != nil {
            panic(err)
        }

        store := storage.NewSQLStore(db)
        b := biz.NewCreateBusiness(store)
        if err := b.CreateProduct(c.Request.Context(), &data); err != nil {
            panic(err)
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(data.Id))
    }
}


