package biz

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
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
    if _, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id}); err != nil {
        if err == common.RecordNotFound {
            return common.ErrEntityNotFound("Product", err)
        }
        return common.ErrCannotGetEntity("Product", err)
    }

    if err := biz.store.UpdateProduct(ctx, id, data); err != nil {
        return common.ErrCannotUpdateEntity("Product", err)
    }
    return nil
}


