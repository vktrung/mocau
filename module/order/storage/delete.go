package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
)

func (s *sqlStore) DeleteOrder(ctx context.Context, id int) error {
	// Soft delete order
	if err := s.db.Delete(&model.Order{}, id).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
