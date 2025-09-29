package ginOrder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/biz"
	"mocau-backend/module/order/model"
	"mocau-backend/module/order/storage"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with customer information and optional order items
// @Tags orders
// @Accept json
// @Produce json
// @Param order body model.OrderCreate true "Order information with optional order items"
// @Success 201 {object} common.Response{data=model.Order}
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /orders [post]
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.OrderCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateOrderBusiness(store)

		if err := business.CreateOrder(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRes(data))
	}
}
