package biz

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/blog/model"
	"mocau-backend/module/upload"
	"strings"
)

type UpdateStorage interface {
	UpdateBlog(ctx context.Context, id int, data *model.BlogUpdate) error
	GetBlog(ctx context.Context, conditions map[string]interface{}) (*model.Blog, error)
}

type updateBusiness struct {
	store UpdateStorage
}

func NewUpdateBusiness(store UpdateStorage) *updateBusiness {
	return &updateBusiness{store: store}
}

func (biz *updateBusiness) UpdateBlog(ctx context.Context, id int, data *model.BlogUpdate) error {
	// Lấy thông tin blog hiện tại để xóa ảnh cũ nếu cần
	currentBlog, err := biz.store.GetBlog(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound("Blog", err)
		}
		return common.ErrCannotGetEntity("Blog", err)
	}

	// Validate title nếu có
	if data.Title != nil {
		*data.Title = strings.TrimSpace(*data.Title)
		if *data.Title == "" {
			return common.ErrInvalidRequest(common.NewCustomError(nil, "blog title cannot be empty", "ErrBlogTitleEmpty"))
		}
	}

	// Validate content nếu có
	if data.Content != nil {
		*data.Content = strings.TrimSpace(*data.Content)
		if *data.Content == "" {
			return common.ErrInvalidRequest(common.NewCustomError(nil, "blog content cannot be empty", "ErrBlogContentEmpty"))
		}
		// Sanitize HTML content để tránh XSS attacks
		*data.Content = common.SanitizeBlogHTML(*data.Content)
	}

	// Validate status nếu có
	if data.Status != nil {
		if *data.Status != "draft" && *data.Status != "published" {
			return common.ErrInvalidRequest(common.NewCustomError(nil, "status must be 'draft' or 'published'", "ErrInvalidStatus"))
		}
	}

	// Nếu có ảnh mới và blog hiện tại có ảnh, xóa ảnh cũ
	if data.Image != nil && currentBlog.Image != nil {
		if err := upload.DeleteImageFromProduct(currentBlog.Image); err != nil {
			// Log error nhưng không dừng quá trình update
		}
	}

	if err := biz.store.UpdateBlog(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity("Blog", err)
	}
	return nil
}
