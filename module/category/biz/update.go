package biz

import (
    "context"
    "mocau-backend/common"
    "mocau-backend/module/category/model"
)

type UpdateStorage interface {
    UpdateCategory(ctx context.Context, id int, data *model.CategoryUpdate) error
    GetCategory(ctx context.Context, conditions map[string]interface{}) (*model.Category, error)
}

type updateBusiness struct {
    store UpdateStorage
}

func NewUpdateBusiness(store UpdateStorage) *updateBusiness {
    return &updateBusiness{store: store}
}

func (biz *updateBusiness) UpdateCategory(ctx context.Context, id int, data *model.CategoryUpdate) error {
    if _, err := biz.store.GetCategory(ctx, map[string]interface{}{"id": id}); err != nil {
        if err == common.RecordNotFound {
            return common.ErrEntityNotFound(model.EntityName, err)
        }
        return common.ErrCannotGetEntity(model.EntityName, err)
    }

    if err := biz.store.UpdateCategory(ctx, id, data); err != nil {
        return common.ErrCannotUpdateEntity(model.EntityName, err)
    }
    return nil
}


