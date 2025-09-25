package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
)

func (s *sqlStore) ListProducts(ctx context.Context) ([]model.Product, error) {
    db := s.db.Table(model.Product{}.TableName())

    var result []model.Product
    if err := db.Find(&result).Error; err != nil {
        return nil, common.ErrDB(err)
    }
    return result, nil
}


