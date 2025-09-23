package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
)

func (s *sqlStore) DeleteCategory(ctx context.Context, id int) error {
    // Soft delete: set status to 'deactive'
    if err := s.db.Table(model.Category{}.TableName()).Where("id = ?", id).Update("status", "deactive").Error; err != nil {
        return common.ErrDB(err)
    }
    return nil
}


