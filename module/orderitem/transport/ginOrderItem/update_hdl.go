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

// UpdateOrderItem godoc
// @Summary Update order item
// @Description Update order item information (only for pending orders)
// @Tags order-items
// @Accept json
// @Produce json
// @Param id path int true "Order Item ID"
// @Param orderItem body model.OrderItemUpdate true "Order item update data"
// @Success 200 {object} common.Response{data=model.OrderItem}
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /order-items/{id} [put]
func UpdateOrderItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data model.OrderItemUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateOrderItemBusiness(store)

		if err := business.UpdateOrderItem(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		// Get updated order item
		getBusiness := biz.NewGetOrderItemBusiness(store)
		updatedOrderItem, err := getBusiness.GetOrderItem(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(updatedOrderItem))
	}
}

// UpdateOrderItemQuantity godoc
// @Summary Update order item quantity
// @Description Update order item quantity (only for pending orders)
// @Tags order-items
// @Accept json
// @Produce json
// @Param id path int true "Order Item ID"
// @Param quantity body map[string]int true "Quantity update data"
// @Success 200 {object} common.Response{data=model.OrderItem}
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /order-items/{id}/quantity [put]
func UpdateOrderItemQuantity(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var request struct {
			Quantity int `json:"quantity" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateOrderItemBusiness(store)

		if err := business.UpdateOrderItemQuantity(c.Request.Context(), id, request.Quantity); err != nil {
			panic(err)
		}

		// Get updated order item
		getBusiness := biz.NewGetOrderItemBusiness(store)
		updatedOrderItem, err := getBusiness.GetOrderItem(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(updatedOrderItem))
	}
}