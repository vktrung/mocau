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

// UpdateProduct godoc
// @Summary Update a product by ID
// @Description Update product fields by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.ProductUpdate true "Product update data"
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
        if err := c.ShouldBind(&data); err != nil {
            panic(err)
        }

        store := storage.NewSQLStore(db)
        b := biz.NewUpdateBusiness(store)
        if err := b.UpdateProduct(c.Request.Context(), id, &data); err != nil {
            panic(err)
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
    }
}


