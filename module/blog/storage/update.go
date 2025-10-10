package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

func (s *sqlStore) UpdateBlog(ctx context.Context, id int, data *model.BlogUpdate) error {
	// Tạo map để chỉ update những trường có giá trị (không phải nil)
	updates := make(map[string]interface{})
	
	if data.Title != nil {
		updates["title"] = *data.Title
	}
	if data.Content != nil {
		updates["content"] = *data.Content
	}
	if data.Status != nil {
		updates["status"] = *data.Status
	}
	if data.Image != nil {
		updates["image"] = data.Image
	}
	
	// Chỉ update nếu có ít nhất một trường cần update
	if len(updates) > 0 {
		if err := s.db.Table(model.Blog{}.TableName()).Where("id = ?", id).Updates(updates).Error; err != nil {
			return common.ErrDB(err)
		}
	}
	
	return nil
}
