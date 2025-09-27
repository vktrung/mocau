package biz

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
)

type GetStorage interface {
	GetBlog(ctx context.Context, conditions map[string]interface{}) (*model.Blog, error)
	GetBlogWithAuthor(ctx context.Context, id int) (*model.BlogWithAuthor, error)
}

type getBusiness struct {
	store GetStorage
}

func NewGetBusiness(store GetStorage) *getBusiness {
	return &getBusiness{store: store}
}

func (biz *getBusiness) GetBlog(ctx context.Context, id int) (*model.Blog, error) {
	result, err := biz.store.GetBlog(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound("Blog", err)
		}
		return nil, common.ErrCannotGetEntity("Blog", err)
	}
	return result, nil
}

func (biz *getBusiness) GetBlogWithAuthor(ctx context.Context, id int) (*model.BlogWithAuthor, error) {
	result, err := biz.store.GetBlogWithAuthor(ctx, id)
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound("Blog", err)
		}
		return nil, common.ErrCannotGetEntity("Blog", err)
	}
	return result, nil
}
