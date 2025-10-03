package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/product/model"
	ordermodel "mocau-backend/module/order/model"
	"time"
)

func (s *sqlStore) GetTopSellingProducts(ctx context.Context, limit int) ([]model.TopSellingProduct, error) {
	// Lấy tháng hiện tại
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	var result []model.TopSellingProduct

	// Query để lấy top sản phẩm bán chạy nhất trong tháng
	// Chỉ tính các order đã completed
	err := s.db.Table("order_items oi").
		Select(`
			oi.product_id,
			p.name as product_name,
			SUM(oi.quantity) as total_sold,
			SUM(oi.quantity * oi.price) as total_revenue,
			p.image
		`).
		Joins("JOIN products p ON oi.product_id = p.id").
		Joins("JOIN orders o ON oi.order_id = o.id").
		Where("o.status = ? AND o.created_at >= ? AND o.created_at <= ?", 
			ordermodel.OrderStatusCompleted, startOfMonth, endOfMonth).
		Group("oi.product_id, p.name, p.image").
		Order("total_sold DESC").
		Limit(limit).
		Scan(&result).Error

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
