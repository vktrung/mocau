package ginOrderItem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/biz"
	"mocau-backend/module/orderitem/model"
	"mocau-backend/module/orderitem/storage"
)

// BulkCreateOrderItems godoc
// @Summary Bulk create order items
// @Description Add multiple items to an order at once
// @Tags order-items
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Param items body []model.OrderItemCreate true "Order items to create"
// @Success 201 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /orders/{order_id}/items/bulk [post]
func BulkCreateOrderItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order_id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var items []model.OrderItemCreate
		if err := c.ShouldBindJSON(&items); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewBulkOrderItemBusiness(store)

		if err := business.BulkCreateOrderItems(c.Request.Context(), orderId, items); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRes(nil))
	}
}

// BulkUpdateOrderItems godoc
// @Summary Bulk update order items
// @Description Update multiple order items at once
// @Tags order-items
// @Accept json
// @Produce json
// @Param updates body map[string]model.OrderItemUpdate true "Order item updates"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /order-items/bulk [put]
func BulkUpdateOrderItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updates map[int]model.OrderItemUpdate
		if err := c.ShouldBindJSON(&updates); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewBulkOrderItemBusiness(store)

		if err := business.BulkUpdateOrderItems(c.Request.Context(), updates); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(nil))
	}
}

// BulkDeleteOrderItems godoc
// @Summary Bulk delete order items
// @Description Delete multiple order items at once
// @Tags order-items
// @Accept json
// @Produce json
// @Param ids body []int true "Order item IDs to delete"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /order-items/bulk [delete]
func BulkDeleteOrderItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ids []int
		if err := c.ShouldBindJSON(&ids); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewBulkOrderItemBusiness(store)

		if err := business.BulkDeleteOrderItems(c.Request.Context(), ids); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(nil))
	}
}