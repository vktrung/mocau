package ginProduct

import (
	"mocau-backend/common"
	"mocau-backend/module/product/biz"
	"mocau-backend/module/product/model"
	"mocau-backend/module/product/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTopSellingProducts godoc
// @Summary Get top selling products
// @Description Get top best selling products of current month based on completed orders. Returns products sorted by total quantity sold.
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int false "Number of products to return (default: 5, max: 20)" minimum(1) maximum(20)
// @Success 200 {object} common.Response{data=[]model.TopSellingProduct} "Successfully retrieved top selling products"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /products/top-selling [get]
func GetTopSellingProducts(db *gorm.DB) func(*gin.Context) {
	// dummy reference for Swagger to resolve model types used in annotations
	var _ = []model.TopSellingProduct{}
	return func(c *gin.Context) {
		// Parse limit parameter, default to 5
		limitStr := c.DefaultQuery("limit", "5")
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			limit = 5
		}

		// Maximum limit to prevent abuse
		if limit > 20 {
			limit = 20
		}

		store := storage.NewSQLStore(db)
		b := biz.NewTopSellingBusiness(store)
		data, err := b.GetTopSellingProducts(c.Request.Context(), limit)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
	}
}

