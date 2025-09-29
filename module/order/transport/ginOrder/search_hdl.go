package ginOrder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/biz"
	"mocau-backend/module/order/storage"
)

// SearchOrders godoc
// @Summary Search orders
// @Description Advanced search for orders with multiple filters
// @Tags orders
// @Accept json
// @Produce json
// @Param query query string false "Search in customer name, phone, order number"
// @Param status query string false "Order status"
// @Param processed_by query int false "Processed by user ID"
// @Param date_from query string false "Date from (YYYY-MM-DD)"
// @Param date_to query string false "Date to (YYYY-MM-DD)"
// @Param min_amount query number false "Minimum amount"
// @Param max_amount query number false "Maximum amount"
// @Success 200 {object} common.Response{data=[]model.Order}
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /orders/search [get]
func SearchOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter storage.OrderSearchFilter

		// Parse query parameters
		if err := c.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewSearchOrderBusiness(store)

		orders, err := business.SearchOrders(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(orders))
	}
}
