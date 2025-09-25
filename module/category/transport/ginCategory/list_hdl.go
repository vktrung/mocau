package ginCategory

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "mocau-backend/common"
    "mocau-backend/module/category/biz"
    "mocau-backend/module/category/storage"
)

// ListCategories godoc
// @Summary List all categories
// @Description Returns all active categories by default.
// @Tags categories
// @Produce json
// @Success 200 {object} common.Response
// @Failure 500 {object} common.AppError "Internal error"
// @Router /categories [get]
func ListCategories(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        store := storage.NewSQLStore(db)
        business := biz.NewListBusiness(store)
        result, err := business.ListCategory(c.Request.Context(), map[string]interface{}{})
        if err != nil {
            c.JSON(http.StatusBadRequest, err)
            return
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(result))
    }
}


