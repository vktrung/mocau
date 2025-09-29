package ginOrder

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/biz"
	"mocau-backend/module/order/storage"
)

// GetOrder godoc
// @Summary Get order by ID
// @Description Get order details by order ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} common.Response{data=model.Order}
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /orders/{id} [get]
func GetOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewGetOrderBusiness(store)

		order, err := business.GetOrder(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(order))
	}
}

// GetOrderByOrderNumber godoc
// @Summary Get order by order number
// @Description Get order details by order number
// @Tags orders
// @Accept json
// @Produce json
// @Param order_number path string true "Order Number"
// @Success 200 {object} common.Response{data=model.Order}
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /orders/number/{order_number} [get]
func GetOrderByOrderNumber(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderNumber := c.Param("order_number")
		if orderNumber == "" {
			panic(common.ErrInvalidRequest(errors.New("order number is required")))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewGetOrderBusiness(store)

		order, err := business.GetOrderByOrderNumber(c.Request.Context(), orderNumber)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(order))
	}
}
