package biz

import (
	"context"
	"mocau-backend/module/user/model"
)

type ListUserStorage interface {
	ListUsers(ctx context.Context, filter *model.UserFilter, moreInfo ...string) ([]model.User, error)
}

type ListUserBusiness struct {
	store ListUserStorage
}

func NewListUserBusiness(store ListUserStorage) *ListUserBusiness {
	return &ListUserBusiness{store: store}
}

func (business *ListUserBusiness) ListUsers(ctx context.Context, filter *model.UserFilter) ([]model.User, error) {
	users, err := business.store.ListUsers(ctx, filter)

	if err != nil {
		return nil, err
	}

	return users, nil
}
