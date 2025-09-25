package biz

import (
    "context"
    "mocau-backend/module/product/model"
)

type ListStorage interface {
    ListProducts(ctx context.Context) ([]model.Product, error)
}

type listBusiness struct {
    store ListStorage
}

func NewListBusiness(store ListStorage) *listBusiness {
    return &listBusiness{store: store}
}

func (biz *listBusiness) ListProducts(ctx context.Context) ([]model.Product, error) {
    return biz.store.ListProducts(ctx)
}


