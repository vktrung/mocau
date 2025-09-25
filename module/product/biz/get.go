package biz

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
)

type GetStorage interface {
    GetProduct(ctx context.Context, conditions map[string]interface{}) (*model.Product, error)
}

type getBusiness struct {
    store GetStorage
}

func NewGetBusiness(store GetStorage) *getBusiness {
    return &getBusiness{store: store}
}

func (biz *getBusiness) GetProduct(ctx context.Context, id int) (*model.Product, error) {
    data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
    if err != nil {
        if err == common.RecordNotFound {
            return nil, common.ErrEntityNotFound("Product", err)
        }
        return nil, common.ErrCannotGetEntity("Product", err)
    }
    return data, nil
}


