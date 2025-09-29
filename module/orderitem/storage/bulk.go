package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/model"
)

func (s *sqlStore) BulkCreateOrderItems(ctx context.Context, orderId int, items []model.OrderItemCreate) error {
	db := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	for _, item := range items {
		item.OrderId = orderId
		if err := s.CreateOrderItem(ctx, &item); err != nil {
			db.Rollback()
			return err
		}
	}

	return db.Commit().Error
}

func (s *sqlStore) BulkUpdateOrderItems(ctx context.Context, updates map[int]model.OrderItemUpdate) error {
	db := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	for id, update := range updates {
		if err := s.UpdateOrderItem(ctx, id, &update); err != nil {
			db.Rollback()
			return err
		}
	}

	return db.Commit().Error
}

func (s *sqlStore) BulkDeleteOrderItems(ctx context.Context, ids []int) error {
	if err := s.db.Delete(&model.OrderItem{}, ids).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
