package biz

import (
	"context"
	"mocau-backend/module/user/model"
)

type UpdateUserProfileStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
	UpdateUserProfile(ctx context.Context, id int, data *model.UserUpdate) error
}

type UpdateUserProfileBusiness struct {
	store UpdateUserProfileStorage
}

func NewUpdateUserProfileBusiness(store UpdateUserProfileStorage) *UpdateUserProfileBusiness {
	return &UpdateUserProfileBusiness{store: store}
}

func (business *UpdateUserProfileBusiness) UpdateUserProfile(ctx context.Context, id int, data *model.UserUpdate) error {
	// Kiểm tra user có tồn tại không
	user, err := business.store.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	// Kiểm tra email có bị trùng không (nếu có update email)
	if data.Email != nil && *data.Email != user.Email {
		existingUser, err := business.store.FindUser(ctx, map[string]interface{}{"email": *data.Email})
		if err == nil && existingUser.Id != id {
			return model.ErrEmailExisted
		}
	}

	// Cập nhật profile
	if err := business.store.UpdateUserProfile(ctx, id, data); err != nil {
		return err
	}

	return nil
}
