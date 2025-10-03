package biz

import (
	"context"
	"mocau-backend/module/product/model"
	"mocau-backend/module/product/storage"
)

type topSellingBusiness struct {
	store storage.Store
}

func NewTopSellingBusiness(store storage.Store) *topSellingBusiness {
	return &topSellingBusiness{store: store}
}

func (b *topSellingBusiness) GetTopSellingProducts(ctx context.Context, limit int) ([]model.TopSellingProduct, error) {
	return b.store.GetTopSellingProducts(ctx, limit)
}

