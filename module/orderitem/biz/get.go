package biz

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/module/orderitem/model"
	"mocau-backend/module/orderitem/storage"
)

type GetOrderItemStorage interface {
	GetDB() *gorm.DB
}

type getOrderItemBusiness struct {
	store GetOrderItemStorage
}

func NewGetOrderItemBusiness(store GetOrderItemStorage) *getOrderItemBusiness {
	return &getOrderItemBusiness{store: store}
}

func (biz *getOrderItemBusiness) GetOrderItem(ctx context.Context, id int) (*model.OrderItem, error) {
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.GetOrderItem(ctx, id)
}

func (biz *getOrderItemBusiness) ListOrderItemsByOrder(ctx context.Context, orderId int) ([]model.OrderItem, error) {
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.ListOrderItemsByOrder(ctx, orderId)
}
