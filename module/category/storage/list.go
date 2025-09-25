package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
)

func (s *sqlStore) ListCategory(ctx context.Context, filter map[string]interface{}) ([]model.Category, error) {
    db := s.db.Table(model.Category{}.TableName())

    var result []model.Category

    if _, ok := filter["status"]; !ok {
        filter["status"] = "active"
    }

    if err := db.Where(filter).Find(&result).Error; err != nil {
        return nil, common.ErrDB(err)
    }

    return result, nil
}


