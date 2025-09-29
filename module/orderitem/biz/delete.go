package biz

import (
	"context"
	"mocau-backend/module/orderitem/storage"
	ordermodel "mocau-backend/module/order/model"
)

type DeleteOrderItemStorage interface {
	GetDB() *gorm.DB
}

type deleteOrderItemBusiness struct {
	store DeleteOrderItemStorage
}

func NewDeleteOrderItemBusiness(store DeleteOrderItemStorage) *deleteOrderItemBusiness {
	return &deleteOrderItemBusiness{store: store}
}

func (biz *deleteOrderItemBusiness) DeleteOrderItem(ctx context.Context, id int) error {
	store := storage.NewSQLStore(biz.store.GetDB())

	// Business Rules Validation

	// 1. Check if order item exists
	orderItem, err := store.GetOrderItem(ctx, id)
	if err != nil {
		return err
	}

	// 2. Check if order is in pending status
	var order ordermodel.Order
	if err := biz.store.GetDB().Where("id = ?", orderItem.OrderId).First(&order).Error; err != nil {
		return ErrOrderNotFound
	}

	// Only allow deleting items from pending orders
	if order.Status != ordermodel.OrderStatusPending {
		return ErrOrderNotPending
	}

	// 3. Delete order item
	return store.DeleteOrderItem(ctx, id)
}
