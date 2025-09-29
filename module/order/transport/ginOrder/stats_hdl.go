package ginOrder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/biz"
	"mocau-backend/module/order/storage"
)

// GetOrderStats godoc
// @Summary Get order statistics
// @Description Get order statistics including counts and revenue
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=storage.OrderStats}
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /orders/stats [get]
func GetOrderStats(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		business := biz.NewGetOrderStatsBusiness(store)

		stats, err := business.GetOrderStats(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(stats))
	}
}
