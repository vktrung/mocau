package biz

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
	"mocau-backend/module/order/storage"
)

type ListOrderStorage interface {
	GetDB() *gorm.DB
}

type listOrderBusiness struct {
	store ListOrderStorage
}

func NewListOrderBusiness(store ListOrderStorage) *listOrderBusiness {
	return &listOrderBusiness{store: store}
}

func (biz *listOrderBusiness) ListOrders(ctx context.Context, filter *model.OrderFilter) ([]model.Order, error) {
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.ListOrders(ctx, filter)
}
