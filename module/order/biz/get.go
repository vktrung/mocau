package biz

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/module/order/model"
	"mocau-backend/module/order/storage"
)

type GetOrderStorage interface {
	GetDB() *gorm.DB
}

type getOrderBusiness struct {
	store GetOrderStorage
}

func NewGetOrderBusiness(store GetOrderStorage) *getOrderBusiness {
	return &getOrderBusiness{store: store}
}

func (biz *getOrderBusiness) GetOrder(ctx context.Context, id int) (*model.Order, error) {
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.GetOrder(ctx, map[string]interface{}{"id": id}, "Processor")
}

func (biz *getOrderBusiness) GetOrderByOrderNumber(ctx context.Context, orderNumber string) (*model.Order, error) {
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.GetOrderByOrderNumber(ctx, orderNumber)
}
