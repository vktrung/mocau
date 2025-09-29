package storage

import (
	"context"
	"mocau-backend/module/order/model"
)

type OrderSearchFilter struct {
	Query         string `json:"query" form:"query"`                   // Search in customer name, phone, order number
	Status        string `json:"status" form:"status"`
	ProcessedBy   int    `json:"processed_by" form:"processed_by"`
	DateFrom      string `json:"date_from" form:"date_from"`           // YYYY-MM-DD
	DateTo        string `json:"date_to" form:"date_to"`               // YYYY-MM-DD
	MinAmount     float64 `json:"min_amount" form:"min_amount"`
	MaxAmount     float64 `json:"max_amount" form:"max_amount"`
}

func (s *sqlStore) SearchOrders(ctx context.Context, filter *OrderSearchFilter) ([]model.Order, error) {
	var orders []model.Order

	db := s.db.Model(&model.Order{})

	// Apply search query
	if filter.Query != "" {
		db = db.Where("(customer_name LIKE ? OR customer_phone LIKE ? OR order_number LIKE ?)",
			"%"+filter.Query+"%", "%"+filter.Query+"%", "%"+filter.Query+"%")
	}

	// Apply filters
	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}
	if filter.ProcessedBy > 0 {
		db = db.Where("processed_by = ?", filter.ProcessedBy)
	}
	if filter.DateFrom != "" {
		db = db.Where("DATE(created_at) >= ?", filter.DateFrom)
	}
	if filter.DateTo != "" {
		db = db.Where("DATE(created_at) <= ?", filter.DateTo)
	}
	if filter.MinAmount > 0 {
		db = db.Where("total_amount >= ?", filter.MinAmount)
	}
	if filter.MaxAmount > 0 {
		db = db.Where("total_amount <= ?", filter.MaxAmount)
	}

	// Preload relations
	db = db.Preload("Processor")

	// Order by created_at desc
	db = db.Order("created_at DESC")

	if err := db.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
