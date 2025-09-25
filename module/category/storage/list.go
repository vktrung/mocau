package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
)

func (s *sqlStore) ListCategory(ctx context.Context, filter map[string]interface{}, paging *common.Paging) ([]model.Category, error) {
    db := s.db.Table(model.Category{}.TableName())

    var result []model.Category

    if _, ok := filter["status"]; !ok {
        filter["status"] = "active"
    }

    if err := db.Where(filter).Count(&paging.Total).Error; err != nil {
        return nil, common.ErrDB(err)
    }

    paging.Process()

    if err := db.Where(filter).
        Offset((paging.Page - 1) * paging.Limit).
        Limit(paging.Limit).
        Find(&result).Error; err != nil {
        return nil, common.ErrDB(err)
    }

    return result, nil
}


