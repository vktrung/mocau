package biz

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
    "strings"
)

type CreateStorage interface {
    CreateCategory(ctx context.Context, data *model.CategoryCreate) error
    GetCategory(ctx context.Context, conditions map[string]interface{}) (*model.Category, error)
}

type createBusiness struct {
    store CreateStorage
}

func NewCreateBusiness(store CreateStorage) *createBusiness {
    return &createBusiness{store: store}
}

func (biz *createBusiness) CreateCategory(ctx context.Context, data *model.CategoryCreate) error {
    data.CategoryName = strings.TrimSpace(data.CategoryName)
    if data.CategoryName == "" {
        return common.ErrInvalidRequest(model.ErrCategoryNameCannotBeEmpty)
    }

    // uniqueness by name (optional)
    if existed, _ := biz.store.GetCategory(ctx, map[string]interface{}{"category_name": data.CategoryName}); existed != nil {
        return common.ErrEntityExisted(model.EntityName, nil)
    }

    if err := biz.store.CreateCategory(ctx, data); err != nil {
        return common.ErrCannotCreateEntity(model.EntityName, err)
    }
    return nil
}


