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

// DeleteOrderItem godoc
// @Summary Delete order item
// @Description Delete order item from order (only for pending orders)
// @Tags order-items
// @Accept json
// @Produce json
// @Param id path int true "Order Item ID"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /order-items/{id} [delete]
func DeleteOrderItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewDeleteOrderItemBusiness(store)

		if err := business.DeleteOrderItem(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(nil))
	}
}