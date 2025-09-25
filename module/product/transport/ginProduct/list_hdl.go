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

// ListProducts godoc
// @Summary List all products
// @Description Retrieve all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=[]model.Product}
// @Router /products [get]
func ListProducts(db *gorm.DB) func(*gin.Context) {
    // dummy reference for Swagger to resolve model types used in annotations
    var _ = []model.Product{}
    return func(c *gin.Context) {
        store := storage.NewSQLStore(db)
        b := biz.NewListBusiness(store)
        data, err := b.ListProducts(c.Request.Context())
        if err != nil {
            panic(err)
        }
        c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
    }
}


