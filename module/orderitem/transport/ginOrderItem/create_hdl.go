package ginOrderItem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/biz"
	"mocau-backend/module/orderitem/model"
	"mocau-backend/module/orderitem/storage"
)

// CreateOrderItem godoc
// @Summary Add item to order
// @Description Add a new item to an existing order (only for pending orders)
// @Tags order-items
// @Accept json
// @Produce json
// @Param orderItem body model.OrderItemCreate true "Order item information"
// @Success 201 {object} common.Response{data=model.OrderItem}
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /orders/{order_id}/items [post]
func CreateOrderItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.OrderItemCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateOrderItemBusiness(store)

		if err := business.CreateOrderItem(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRes(data))
	}
}
