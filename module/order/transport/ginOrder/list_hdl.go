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

// ListOrders godoc
// @Summary List orders
// @Description Get list of orders with filtering
// @Tags orders
// @Accept json
// @Produce json
// @Param status query string false "Order status (pending, confirmed, completed, cancelled)"
// @Param processed_by query int false "Processed by user ID"
// @Param customer_phone query string false "Customer phone number"
// @Param order_number query string false "Order number"
// @Success 200 {object} common.Response{data=[]model.Order}
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /orders [get]
func ListOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.OrderFilter

		// Parse query parameters
		if err := c.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewListOrderBusiness(store)

		orders, err := business.ListOrders(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(orders))
	}
}
