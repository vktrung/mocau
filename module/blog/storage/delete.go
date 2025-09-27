package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

func (s *sqlStore) DeleteBlog(ctx context.Context, id int) error {
	if err := s.db.Table(model.Blog{}.TableName()).Where("id = ?", id).Update("deleted_at", common.GetCurrentTime()).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
