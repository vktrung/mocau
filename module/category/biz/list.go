package biz

import (
    "context"
    "mocau-backend/module/category/model"
)

type ListStorage interface {
    ListCategory(ctx context.Context, filter map[string]interface{}) ([]model.Category, error)
}

type listBusiness struct {
    store ListStorage
}

func NewListBusiness(store ListStorage) *listBusiness {
    return &listBusiness{store: store}
}

func (biz *listBusiness) ListCategory(ctx context.Context, filter map[string]interface{}) ([]model.Category, error) {
    return biz.store.ListCategory(ctx, filter)
}


