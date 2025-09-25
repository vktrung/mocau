package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
)

func (s *sqlStore) CreateProduct(ctx context.Context, data *model.ProductCreate) error {
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


