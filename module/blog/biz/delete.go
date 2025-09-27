package biz

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

type DeleteStorage interface {
	DeleteBlog(ctx context.Context, id int) error
	GetBlog(ctx context.Context, conditions map[string]interface{}) (*model.Blog, error)
}

type deleteBusiness struct {
	store DeleteStorage
}

func NewDeleteBusiness(store DeleteStorage) *deleteBusiness {
	return &deleteBusiness{store: store}
}

func (biz *deleteBusiness) DeleteBlog(ctx context.Context, id int) error {
	// Kiểm tra blog có tồn tại không
	if _, err := biz.store.GetBlog(ctx, map[string]interface{}{"id": id}); err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound("Blog", err)
		}
		return common.ErrCannotGetEntity("Blog", err)
	}

	if err := biz.store.DeleteBlog(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity("Blog", err)
	}
	return nil
}
