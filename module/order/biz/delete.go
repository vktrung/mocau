package biz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
	"mocau-backend/module/order/storage"
)

type DeleteOrderStorage interface {
	GetDB() *gorm.DB
}

type deleteOrderBusiness struct {
	store DeleteOrderStorage
}

func NewDeleteOrderBusiness(store DeleteOrderStorage) *deleteOrderBusiness {
	return &deleteOrderBusiness{store: store}
}

func (biz *deleteOrderBusiness) DeleteOrder(ctx context.Context, id int) error {
	store := storage.NewSQLStore(biz.store.GetDB())

	// Business Rules Validation

	// 1. Check if order exists
	order, err := store.GetOrder(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	// 2. Only allow deleting pending orders
	if order.Status != model.OrderStatusPending {
		return ErrCannotDeleteNonPendingOrder
	}

	// 3. Delete order
	return store.DeleteOrder(ctx, id)
}

// Business Rule Errors
var (
	ErrCannotDeleteNonPendingOrder = common.NewCustomError(errors.New("cannot delete non-pending order"), "can only delete pending orders", "ErrCannotDeleteNonPendingOrder")
)
