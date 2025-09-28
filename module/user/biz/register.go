package biz

import (
	"mocau-backend/common"
	"mocau-backend/module/user/model"
	"context"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(regiterStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: regiterStorage,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *model.UserCreate) error {
	// Check if email already exists
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return model.ErrEmailExisted
	}

	// Check if username already exists
	user, _ = business.registerStorage.FindUser(ctx, map[string]interface{}{"username": data.Username})
	if user != nil {
		return model.ErrUsernameExisted
	}

	// Hash password (without salt as per new schema)
	data.Password = business.hasher.Hash(data.Password)
	data.Role = "user" // hard code
	data.Status = "active" // set default status

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
