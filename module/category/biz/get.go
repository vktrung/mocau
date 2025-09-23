package biz

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
)

type GetStorage interface {
    GetCategory(ctx context.Context, conditions map[string]interface{}) (*model.Category, error)
}

type getBusiness struct {
    store GetStorage
}

func NewGetBusiness(store GetStorage) *getBusiness {
    return &getBusiness{store: store}
}

func (biz *getBusiness) GetCategory(ctx context.Context, id int) (*model.Category, error) {
    data, err := biz.store.GetCategory(ctx, map[string]interface{}{"id": id})
    if err != nil {
        if err == common.RecordNotFound {
            return nil, common.ErrEntityNotFound(model.EntityName, err)
        }
        return nil, common.ErrCannotGetEntity(model.EntityName, err)
    }
    return data, nil
}


