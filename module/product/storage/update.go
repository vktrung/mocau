package storage

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
)

func (s *sqlStore) UpdateProduct(ctx context.Context, id int, data *model.ProductUpdate) error {
    // Sử dụng Updates thay vì Save để chỉ update các field được set
    result := s.db.Table(model.Product{}.TableName()).Where("id = ?", id).Updates(data)
    if result.Error != nil {
        return common.ErrDB(result.Error)
    }
    
    // Kiểm tra xem có record nào được update không
    if result.RowsAffected == 0 {
        return common.RecordNotFound
    }
    
    return nil
}


