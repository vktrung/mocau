package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
	"time"
)

func (s *sqlStore) DeleteBlog(ctx context.Context, id int) error {
	if err := s.db.Table(model.Blog{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
