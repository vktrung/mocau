package ginOrder

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/biz"
	"mocau-backend/module/order/model"
	"mocau-backend/module/order/storage"
)

// UpdateOrder godoc
// @Summary Update order
// @Description Update order information
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body model.OrderUpdate true "Order update data"
// @Success 200 {object} common.Response{data=model.Order}
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /orders/{id} [put]
func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data model.OrderUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateOrderBusiness(store)

		if err := business.UpdateOrder(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		// Get updated order
		getBusiness := biz.NewGetOrderBusiness(store)
		updatedOrder, err := getBusiness.GetOrder(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(updatedOrder))
	}
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update order status (pending -> confirmed -> completed)
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body object{status=string} true "Status update data"
// @Success 200 {object} common.Response{data=model.Order}
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /orders/{id}/status [put]
func UpdateOrderStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var request struct {
			Status model.OrderStatus `json:"status" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateOrderBusiness(store)

		if err := business.UpdateOrderStatus(c.Request.Context(), id, request.Status, nil); err != nil {
			panic(err)
		}

		// Get updated order
		getBusiness := biz.NewGetOrderBusiness(store)
		updatedOrder, err := getBusiness.GetOrder(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(updatedOrder))
	}
}
