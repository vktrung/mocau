package storage

import (
	"mocau-backend/common"
	"mocau-backend/module/user/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error) {
	db := s.db.Table(model.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user model.User
	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	// Mask the user ID for security
	user.Mask(common.DbTypeUser)

	return &user, nil
}
