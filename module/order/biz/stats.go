package biz

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/module/order/storage"
)

type GetOrderStatsStorage interface {
	GetDB() *gorm.DB
}

type getOrderStatsBusiness struct {
	store GetOrderStatsStorage
}

func NewGetOrderStatsBusiness(store GetOrderStatsStorage) *getOrderStatsBusiness {
	return &getOrderStatsBusiness{store: store}
}

func (biz *getOrderStatsBusiness) GetOrderStats(ctx context.Context) (*storage.OrderStats, error) {
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.GetOrderStats(ctx)
}
