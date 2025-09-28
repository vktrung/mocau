package biz

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/component/tokenprovider"
	"mocau-backend/module/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *model.UserLogin) (tokenprovider.Token, error) {
	var user *model.User
	var err error

	// Try to login with username first, then email
	if data.Username != "" {
		user, err = business.storeUser.FindUser(ctx, map[string]interface{}{"username": data.Username})
	} else if data.Email != "" {
		user, err = business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	} else {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	if err != nil {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	// Hash the provided password and compare
	passHashed := business.hasher.Hash(data.Password)

	if user.Password != passHashed {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	// Check if account is active
	if user.Status != "active" {
		return nil, model.ErrAccountInactive
	}

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
