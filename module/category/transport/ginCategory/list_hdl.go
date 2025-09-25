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
// @Summary List categories with paging
// @Description Returns active categories by default. You can include `status` filter in future if needed.
// @Tags categories
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Page size (min 5, max 100)"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.AppError "Internal error"
// @Router /categories [get]
func ListCategories(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var paging common.Paging
        if err := c.ShouldBindQuery(&paging); err != nil {
            c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
            return
        }

        store := storage.NewSQLStore(db)
        business := biz.NewListBusiness(store)
        result, err := business.ListCategory(c.Request.Context(), map[string]interface{}{}, &paging)
        if err != nil {
            c.JSON(http.StatusBadRequest, err)
            return
        }

        c.JSON(http.StatusOK, common.NewSuccessRes(result, paging, nil))
    }
}


