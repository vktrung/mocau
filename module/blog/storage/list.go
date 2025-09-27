package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

func (s *sqlStore) ListBlogs(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) ([]model.Blog, error) {
	var result []model.Blog

	db := s.db.Table(model.Blog{}.TableName())

	if conditions != nil {
		db = db.Where(conditions)
	}

	for i := range moreKeys {
		db = db.Where(moreKeys[i])
	}

	if err := db.Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
