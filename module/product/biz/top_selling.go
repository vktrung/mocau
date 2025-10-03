package biz

import (
	"context"
	"mocau-backend/module/product/model"
)

type TopSellingStorage interface {
	GetTopSellingProducts(ctx context.Context, limit int) ([]model.TopSellingProduct, error)
}

type topSellingBusiness struct {
	store TopSellingStorage
}

func NewTopSellingBusiness(store TopSellingStorage) *topSellingBusiness {
	return &topSellingBusiness{store: store}
}

func (b *topSellingBusiness) GetTopSellingProducts(ctx context.Context, limit int) ([]model.TopSellingProduct, error) {
	return b.store.GetTopSellingProducts(ctx, limit)
}

