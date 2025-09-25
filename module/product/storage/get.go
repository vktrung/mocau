package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
    "gorm.io/gorm"
)

func (s *sqlStore) GetProduct(ctx context.Context, conditions map[string]interface{}) (*model.Product, error) {
    db := s.db.Table(model.Product{}.TableName())

    var data model.Product
    if err := db.Where(conditions).First(&data).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, common.RecordNotFound
        }
        return nil, common.ErrDB(err)
    }

    return &data, nil
}


