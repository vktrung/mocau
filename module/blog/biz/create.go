package biz

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
	"strings"
)

type CreateStorage interface {
	CreateBlog(ctx context.Context, data *model.BlogCreate) error
	GetBlog(ctx context.Context, conditions map[string]interface{}) (*model.Blog, error)
}

type createBusiness struct {
	store CreateStorage
}

func NewCreateBusiness(store CreateStorage) *createBusiness {
	return &createBusiness{store: store}
}

func (biz *createBusiness) CreateBlog(ctx context.Context, data *model.BlogCreate) error {
	data.Title = strings.TrimSpace(data.Title)
	if data.Title == "" {
		return common.ErrInvalidRequest(common.NewCustomError(nil, "blog title cannot be empty", "ErrBlogTitleEmpty"))
	}

	data.Content = strings.TrimSpace(data.Content)
	if data.Content == "" {
		return common.ErrInvalidRequest(common.NewCustomError(nil, "blog content cannot be empty", "ErrBlogContentEmpty"))
	}

	// Sanitize HTML content để tránh XSS attacks
	data.Content = common.SanitizeBlogHTML(data.Content)

	// Validate status
	if data.Status != "" && data.Status != "draft" && data.Status != "published" {
		return common.ErrInvalidRequest(common.NewCustomError(nil, "status must be 'draft' or 'published'", "ErrInvalidStatus"))
	}

	if err := biz.store.CreateBlog(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("Blog", err)
	}
	return nil
}
