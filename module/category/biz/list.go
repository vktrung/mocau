package biz

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
)

type ListStorage interface {
    ListCategory(ctx context.Context, filter map[string]interface{}, paging *common.Paging) ([]model.Category, error)
}

type listBusiness struct {
    store ListStorage
}

func NewListBusiness(store ListStorage) *listBusiness {
    return &listBusiness{store: store}
}

func (biz *listBusiness) ListCategory(ctx context.Context, filter map[string]interface{}, paging *common.Paging) ([]model.Category, error) {
    items, err := biz.store.ListCategory(ctx, filter, paging)
    if err != nil {
        return nil, common.ErrCannotListEntity(model.EntityName, err)
    }
    return items, nil
}


