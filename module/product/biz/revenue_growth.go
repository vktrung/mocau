package biz

import (
	"context"
	"mocau-backend/module/product/model"
)

type RevenueGrowthStorage interface {
	GetRevenueGrowth(ctx context.Context) (*model.RevenueGrowth, error)
}

type revenueGrowthBusiness struct {
	store RevenueGrowthStorage
}

func NewRevenueGrowthBusiness(store RevenueGrowthStorage) *revenueGrowthBusiness {
	return &revenueGrowthBusiness{store: store}
}

func (b *revenueGrowthBusiness) GetRevenueGrowth(ctx context.Context) (*model.RevenueGrowth, error) {
	return b.store.GetRevenueGrowth(ctx)
}
