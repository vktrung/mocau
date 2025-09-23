package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *model.CategoryCreate) error {
    db := s.db.Begin()

    if err := db.Table(data.TableName()).Create(data).Error; err != nil {
        db.Rollback()
        return common.ErrDB(err)
    }

    if err := db.Commit().Error; err != nil {
        db.Rollback()
        return common.ErrDB(err)
    }

    return nil
}


