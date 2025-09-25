package ginProduct

import (
    "mocau-backend/common"
    "mocau-backend/module/product/biz"
    "mocau-backend/module/product/model"
    "mocau-backend/module/product/storage"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// GetProduct godoc
// @Summary Get product by ID
// @Description Retrieve product details by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} common.Response{data=model.Product}
// @Failure 404 {object} common.Response "Product not found"
// @Router /products/{id} [get]
func GetProduct(db *gorm.DB) func(*gin.Context) {
    // dummy reference for Swagger to resolve model types used in annotations
    var _ = model.Product{}
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            panic(common.ErrInvalidRequest(err))
        }

        store := storage.NewSQLStore(db)
        b := biz.NewGetBusiness(store)
        data, err := b.GetProduct(c.Request.Context(), id)
        if err != nil {
            panic(err)
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
    }
}


