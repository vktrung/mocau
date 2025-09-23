package ginCategory

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "mocau-backend/common"
    "mocau-backend/module/category/biz"
    "mocau-backend/module/category/model"
    "mocau-backend/module/category/storage"
)

// CreateCategory godoc
// @Summary Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param data body model.CategoryCreate true "Category data"
// @Success 201 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /categories [post]
func CreateCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var data model.CategoryCreate
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
            return
        }

        store := storage.NewSQLStore(db)
        business := biz.NewCreateBusiness(store)
        if err := business.CreateCategory(c.Request.Context(), &data); err != nil {
            c.JSON(http.StatusBadRequest, err)
            return
        }

        c.JSON(http.StatusCreated, common.SimpleSuccessRes(true))
    }
}


