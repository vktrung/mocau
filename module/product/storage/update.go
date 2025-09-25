package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
)

func (s *sqlStore) UpdateProduct(ctx context.Context, id int, data *model.ProductUpdate) error {
    if err := s.db.Table(model.Product{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
        return common.ErrDB(err)
    }
    return nil
}


