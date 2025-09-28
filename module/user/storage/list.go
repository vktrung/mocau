package storage

import (
	"mocau-backend/common"
	"mocau-backend/module/user/model"
	"context"
)

func (s *sqlStore) ListUsers(ctx context.Context, filter *model.UserFilter, moreInfo ...string) ([]model.User, error) {
	db := s.db.Table(model.User{}.TableName())

	if f := filter; f != nil {
		if f.Status != "" {
			db = db.Where("status = ?", f.Status)
		}
		if f.Role != "" {
			db = db.Where("role = ?", f.Role)
		}
	}

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var users []model.User

	if err := db.Order("id desc").Find(&users).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return users, nil
}
