package biz

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

type ListStorage interface {
	ListBlogs(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) ([]model.Blog, error)
}

type listBusiness struct {
	store ListStorage
}

func NewListBusiness(store ListStorage) *listBusiness {
	return &listBusiness{store: store}
}

func (biz *listBusiness) ListBlogs(ctx context.Context, filter map[string]interface{}) ([]model.Blog, error) {
	result, err := biz.store.ListBlogs(ctx, filter)
	if err != nil {
		return nil, common.ErrCannotListEntity("Blog", err)
	}
	return result, nil
}
