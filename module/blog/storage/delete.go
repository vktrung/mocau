package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
	"time"
)

func (s *sqlStore) DeleteBlog(ctx context.Context, id int) error {
	now := time.Now()
	if err := s.db.Table(model.Blog{}.TableName()).Where("id = ?", id).Update("deleted_at", &now).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
