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

// DeleteCategory godoc
// @Summary Soft-delete a category by id
// @Description Set status to `deactive` instead of removing the record.
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 404 {object} common.AppError
// @Router /categories/{id} [delete]
func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, _ := strconv.Atoi(idStr)

        store := storage.NewSQLStore(db)
        business := biz.NewDeleteBusiness(store)
        if err := business.DeleteCategory(c.Request.Context(), id); err != nil {
            c.JSON(http.StatusBadRequest, err)
            return
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
    }
}


