package biz

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/user/model"
)

type UpdateUserStatusStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
	UpdateUserStatus(ctx context.Context, id int, status string) error
}

type UpdateUserStatusBusiness struct {
	store UpdateUserStatusStorage
}

func NewUpdateUserStatusBusiness(store UpdateUserStatusStorage) *UpdateUserStatusBusiness {
	return &UpdateUserStatusBusiness{store: store}
}

func (business *UpdateUserStatusBusiness) ToggleUserStatus(ctx context.Context, id int) error {
	// Kiểm tra user có tồn tại không
	user, err := business.store.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	// Tự động toggle status
	var newStatus string
	if user.Status == "active" {
		newStatus = "inactive"
	} else {
		newStatus = "active"
	}

	// Cập nhật status
	if err := business.store.UpdateUserStatus(ctx, id, newStatus); err != nil {
		return err
	}

	return nil
}
