package ginCategory

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "mocau-backend/common"
    "mocau-backend/module/category/biz"
    "mocau-backend/module/category/storage"
)

// GetCategory godoc
// @Summary Get category by id
// @Description Return a single active category by numeric id. Soft-deleted categories are excluded.
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} common.Response
// @Failure 404 {object} common.AppError "Not found"
// @Failure 500 {object} common.AppError "Internal error"
// @Router /categories/{id} [get]
func GetCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, _ := strconv.Atoi(idStr)

        store := storage.NewSQLStore(db)
        business := biz.NewGetBusiness(store)
        data, err := business.GetCategory(c.Request.Context(), id)
        if err != nil {
            c.JSON(http.StatusBadRequest, err)
            return
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
    }
}


