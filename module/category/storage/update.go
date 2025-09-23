package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
)

func (s *sqlStore) UpdateCategory(ctx context.Context, id int, data *model.CategoryUpdate) error {
    if err := s.db.Table(model.Category{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
        return common.ErrDB(err)
    }
    return nil
}


