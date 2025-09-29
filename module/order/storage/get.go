package storage

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
)

func (s *sqlStore) GetOrder(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.Order, error) {
	var order model.Order

	db := s.db

	// Preload relations if needed
	for _, info := range moreInfo {
		db = db.Preload(info)
	}

	if err := db.Where(conditions).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &order, nil
}

func (s *sqlStore) GetOrderByOrderNumber(ctx context.Context, orderNumber string) (*model.Order, error) {
	return s.GetOrder(ctx, map[string]interface{}{"order_number": orderNumber}, "Processor")
}
