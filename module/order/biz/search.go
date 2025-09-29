package biz

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/module/order/model"
	"mocau-backend/module/order/storage"
)

type SearchOrderStorage interface {
	GetDB() *gorm.DB
}

type searchOrderBusiness struct {
	store SearchOrderStorage
}

func NewSearchOrderBusiness(store SearchOrderStorage) *searchOrderBusiness {
	return &searchOrderBusiness{store: store}
}

func (biz *searchOrderBusiness) SearchOrders(ctx context.Context, filter *storage.OrderSearchFilter) ([]model.Order, error) {
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.SearchOrders(ctx, filter)
}
