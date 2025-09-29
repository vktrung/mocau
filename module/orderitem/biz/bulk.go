package biz

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/module/orderitem/model"
	"mocau-backend/module/orderitem/storage"
	ordermodel "mocau-backend/module/order/model"
)

type BulkOrderItemStorage interface {
	GetDB() *gorm.DB
}

type bulkOrderItemBusiness struct {
	store BulkOrderItemStorage
}

func NewBulkOrderItemBusiness(store BulkOrderItemStorage) *bulkOrderItemBusiness {
	return &bulkOrderItemBusiness{store: store}
}

func (biz *bulkOrderItemBusiness) BulkCreateOrderItems(ctx context.Context, orderId int, items []model.OrderItemCreate) error {
	// Business Rules Validation

	// 1. Check if order exists and is pending
	var order ordermodel.Order
	if err := biz.store.GetDB().Where("id = ?", orderId).First(&order).Error; err != nil {
		return ErrOrderNotFound
	}

	if order.Status != ordermodel.OrderStatusPending {
		return ErrOrderNotPending
	}

	// 2. Validate all items
	for _, item := range items {
		if item.ProductId <= 0 {
			return ErrInvalidProductId
		}
		if item.Quantity <= 0 {
			return ErrInvalidQuantity
		}
	}

	// 3. Bulk create through storage
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.BulkCreateOrderItems(ctx, orderId, items)
}

func (biz *bulkOrderItemBusiness) BulkUpdateOrderItems(ctx context.Context, updates map[int]model.OrderItemUpdate) error {
	// Business Rules Validation

	// 1. Validate all order items exist and are in pending orders
	for id := range updates {
		orderItem, err := storage.NewSQLStore(biz.store.GetDB()).GetOrderItem(ctx, id)
		if err != nil {
			return err
		}

		var order ordermodel.Order
		if err := biz.store.GetDB().Where("id = ?", orderItem.OrderId).First(&order).Error; err != nil {
			return ErrOrderNotFound
		}

		if order.Status != ordermodel.OrderStatusPending {
			return ErrOrderNotPending
		}
	}

	// 2. Bulk update through storage
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.BulkUpdateOrderItems(ctx, updates)
}

func (biz *bulkOrderItemBusiness) BulkDeleteOrderItems(ctx context.Context, ids []int) error {
	// Business Rules Validation

	// 1. Validate all order items exist and are in pending orders
	for _, id := range ids {
		orderItem, err := storage.NewSQLStore(biz.store.GetDB()).GetOrderItem(ctx, id)
		if err != nil {
			return err
		}

		var order ordermodel.Order
		if err := biz.store.GetDB().Where("id = ?", orderItem.OrderId).First(&order).Error; err != nil {
			return ErrOrderNotFound
		}

		if order.Status != ordermodel.OrderStatusPending {
			return ErrOrderNotPending
		}
	}

	// 2. Bulk delete through storage
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.BulkDeleteOrderItems(ctx, ids)
}
