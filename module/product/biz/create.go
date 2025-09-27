package biz

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/product/model"
    "strings"
)

type CreateStorage interface {
    CreateProduct(ctx context.Context, data *model.ProductCreate) error
    GetProduct(ctx context.Context, conditions map[string]interface{}) (*model.Product, error)
}

type createBusiness struct {
    store CreateStorage
}

func NewCreateBusiness(store CreateStorage) *createBusiness {
    return &createBusiness{store: store}
}

func (biz *createBusiness) CreateProduct(ctx context.Context, data *model.ProductCreate) error {
    data.Name = strings.TrimSpace(data.Name)
    if data.Name == "" {
        return common.ErrInvalidRequest(common.NewCustomError(nil, "product name cannot be empty", "ErrProductNameEmpty"))
    }

    // Image đã được xử lý trong upload service, không cần xử lý thêm ở đây

    if err := biz.store.CreateProduct(ctx, data); err != nil {
        return common.ErrCannotCreateEntity("Product", err)
    }
    return nil
}


