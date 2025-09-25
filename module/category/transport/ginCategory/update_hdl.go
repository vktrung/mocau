package ginCategory

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "mocau-backend/common"
    "mocau-backend/module/category/biz"
    "mocau-backend/module/category/model"
    "mocau-backend/module/category/storage"
)

// UpdateCategory godoc
// @Summary Update a category
// @Description Partial update for category fields. Name must remain unique.
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param data body model.CategoryUpdate true "Category update payload"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError "Invalid payload"
// @Failure 404 {object} common.AppError "Not found"
// @Router /categories/{id} [put]
func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, _ := strconv.Atoi(idStr)

        var data model.CategoryUpdate
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
            return
        }

        store := storage.NewSQLStore(db)
        business := biz.NewUpdateBusiness(store)
        if err := business.UpdateCategory(c.Request.Context(), id, &data); err != nil {
            c.JSON(http.StatusBadRequest, err)
            return
        }

        c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
    }
}


