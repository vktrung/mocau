package biz

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
    "mocau-backend/module/upload"
)

type UpdateStorage interface {
    UpdateProduct(ctx context.Context, id int, data *model.ProductUpdate) error
    GetProduct(ctx context.Context, conditions map[string]interface{}) (*model.Product, error)
}

type updateBusiness struct {
    store UpdateStorage
}

func NewUpdateBusiness(store UpdateStorage) *updateBusiness {
    return &updateBusiness{store: store}
}

func (biz *updateBusiness) UpdateProduct(ctx context.Context, id int, data *model.ProductUpdate) error {
    // Lấy thông tin sản phẩm hiện tại để xóa ảnh cũ nếu cần
    currentProduct, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
    if err != nil {
        if err == common.RecordNotFound {
            return common.ErrEntityNotFound("Product", err)
        }
        return common.ErrCannotGetEntity("Product", err)
    }

    // Nếu có ảnh mới và sản phẩm hiện tại có ảnh, xóa ảnh cũ
    if data.Image != nil && currentProduct.Image != nil {
        if err := upload.DeleteImageFromProduct(currentProduct.Image); err != nil {
            // Log error nhưng không dừng quá trình update
            // Có thể thêm logging ở đây
        }
    }

    // Image đã được xử lý trong upload service, không cần xử lý thêm ở đây

    if err := biz.store.UpdateProduct(ctx, id, data); err != nil {
        return common.ErrCannotUpdateEntity("Product", err)
    }
    return nil
}


