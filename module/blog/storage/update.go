package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

func (s *sqlStore) UpdateBlog(ctx context.Context, id int, data *model.BlogUpdate) error {
	if err := s.db.Table(model.Blog{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
