package ginOrderItem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/biz"
	"mocau-backend/module/orderitem/storage"
)

// GetOrderItem godoc
// @Summary Get order item by ID
// @Description Get order item details by ID
// @Tags order-items
// @Accept json
// @Produce json
// @Param id path int true "Order Item ID"
// @Success 200 {object} common.Response{data=model.OrderItem}
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /order-items/{id} [get]
func GetOrderItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewGetOrderItemBusiness(store)

		orderItem, err := business.GetOrderItem(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(orderItem))
	}
}

// ListOrderItemsByOrder godoc
// @Summary List order items by order ID
// @Description Get all items in a specific order
// @Tags order-items
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} common.Response{data=[]model.OrderItem}
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /order-items/order/{order_id} [get]
func ListOrderItemsByOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order_id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewGetOrderItemBusiness(store)

		orderItems, err := business.ListOrderItemsByOrder(c.Request.Context(), orderId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(orderItems))
	}
}