package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/user/model"
)

func (s *sqlStore) UpdateUserStatus(ctx context.Context, id int, status string) error {
	db := s.db.Table(model.User{}.TableName())

	if err := db.Where("id = ?", id).Update("status", status).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateUserProfile(ctx context.Context, id int, data *model.UserUpdate) error {
	db := s.db.Table(model.User{}.TableName())

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
