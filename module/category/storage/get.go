package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
    "gorm.io/gorm"
)

func (s *sqlStore) GetCategory(ctx context.Context, conditions map[string]interface{}) (*model.Category, error) {
    db := s.db.Table(model.Category{}.TableName())

    if _, ok := conditions["status"]; !ok {
        conditions["status"] = "active"
    }

    var data model.Category
    if err := db.Where(conditions).First(&data).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, common.RecordNotFound
        }
        return nil, common.ErrDB(err)
    }

    data.Mask(common.DbTypeUser)
    return &data, nil
}


