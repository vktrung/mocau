package ginOrder

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/biz"
	"mocau-backend/module/order/storage"
)

// DeleteOrder godoc
// @Summary Delete order
// @Description Delete order (only pending orders can be deleted)
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 500 {object} common.Response
// @Security BearerAuth
// @Router /orders/{id} [delete]
func DeleteOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		business := biz.NewDeleteOrderBusiness(store)

		if err := business.DeleteOrder(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(nil))
	}
}
