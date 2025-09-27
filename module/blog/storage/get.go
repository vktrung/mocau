package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

func (s *sqlStore) GetBlog(ctx context.Context, conditions map[string]interface{}) (*model.Blog, error) {
	var result model.Blog

	if err := s.db.Where(conditions).First(&result).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
