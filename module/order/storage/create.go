package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
	"time"
)

func (s *sqlStore) CreateOrder(ctx context.Context, data *model.OrderCreate) error {
	db := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	// Generate order number
	orderNumber := s.generateOrderNumber()
	
	// Create order
	order := &model.Order{
		OrderNumber:     orderNumber,
		Status:          model.OrderStatusPending,
		TotalAmount:     data.TotalAmount,
		CustomerName:    data.CustomerName,
		CustomerPhone:   data.CustomerPhone,
		CustomerEmail:   data.CustomerEmail,
		ShippingAddress: data.ShippingAddress,
		Notes:           data.Notes,
	}

	if err := db.Create(order).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	// Order items will be created separately through OrderItem APIs

	return db.Commit().Error
}

func (s *sqlStore) generateOrderNumber() string {
	// Generate order number: ORD + timestamp + random
	timestamp := time.Now().Format("20060102150405")
	return "ORD" + timestamp
}
