package ginProduct

import (
	"mocau-backend/common"
	"mocau-backend/module/product/biz"
	"mocau-backend/module/product/model"
	"mocau-backend/module/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetRevenueGrowth godoc
// @Summary Get revenue growth statistics
// @Description Get revenue growth comparison between current month and previous month based on completed orders
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=model.RevenueGrowth} "Successfully retrieved revenue growth statistics"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /products/revenue-growth [get]
func GetRevenueGrowth(db *gorm.DB) func(*gin.Context) {
	// dummy reference for Swagger to resolve model types used in annotations
	var _ = model.RevenueGrowth{}
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		b := biz.NewRevenueGrowthBusiness(store)
		data, err := b.GetRevenueGrowth(c.Request.Context())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}
