package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
)

func (s *sqlStore) ListOrders(ctx context.Context, filter *model.OrderFilter) ([]model.Order, error) {
	var orders []model.Order

	db := s.db.Model(&model.Order{})

	// Apply filters
	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}
	if filter.ProcessedBy > 0 {
		db = db.Where("processed_by = ?", filter.ProcessedBy)
	}
	if filter.CustomerPhone != "" {
		db = db.Where("customer_phone LIKE ?", "%"+filter.CustomerPhone+"%")
	}
	if filter.OrderNumber != "" {
		db = db.Where("order_number LIKE ?", "%"+filter.OrderNumber+"%")
	}

	// Preload relations
	db = db.Preload("Processor")

	// Order by created_at desc
	db = db.Order("created_at DESC")

	if err := db.Find(&orders).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orders, nil
}
