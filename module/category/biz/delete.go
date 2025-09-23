package biz

import (
    "context"
    "mocau-backend/common"
)

type DeleteStorage interface {
    DeleteCategory(ctx context.Context, id int) error
}

type deleteBusiness struct {
    store DeleteStorage
}

func NewDeleteBusiness(store DeleteStorage) *deleteBusiness {
    return &deleteBusiness{store: store}
}

func (biz *deleteBusiness) DeleteCategory(ctx context.Context, id int) error {
    if err := biz.store.DeleteCategory(ctx, id); err != nil {
        return common.ErrCannotDeleteEntity("Category", err)
    }
    return nil
}


