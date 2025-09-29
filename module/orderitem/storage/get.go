package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/model"
	"gorm.io/gorm"
)

func (s *sqlStore) GetOrderItem(ctx context.Context, id int) (*model.OrderItem, error) {
	var orderItem model.OrderItem

	if err := s.db.Preload("Order").Preload("Product").Where("id = ?", id).First(&orderItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &orderItem, nil
}

func (s *sqlStore) ListOrderItemsByOrder(ctx context.Context, orderId int) ([]model.OrderItem, error) {
	var orderItems []model.OrderItem

	if err := s.db.Preload("Product").Where("order_id = ?", orderId).Find(&orderItems).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orderItems, nil
}
