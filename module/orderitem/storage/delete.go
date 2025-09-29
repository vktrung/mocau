package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/model"
)

func (s *sqlStore) DeleteOrderItem(ctx context.Context, id int) error {
	if err := s.db.Delete(&model.OrderItem{}, id).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
