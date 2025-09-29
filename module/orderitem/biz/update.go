package biz

import (
	"context"
	"mocau-backend/module/orderitem/model"
	"mocau-backend/module/orderitem/storage"
	ordermodel "mocau-backend/module/order/model"
	productmodel "mocau-backend/module/product/model"
)

type UpdateOrderItemStorage interface {
	GetDB() *gorm.DB
}

type updateOrderItemBusiness struct {
	store UpdateOrderItemStorage
}

func NewUpdateOrderItemBusiness(store UpdateOrderItemStorage) *updateOrderItemBusiness {
	return &updateOrderItemBusiness{store: store}
}

func (biz *updateOrderItemBusiness) UpdateOrderItem(ctx context.Context, id int, data *model.OrderItemUpdate) error {
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

	// Only allow updating items in pending orders
	if order.Status != ordermodel.OrderStatusPending {
		return ErrOrderNotPending
	}

	// 3. Validate quantity if provided
	if data.Quantity != nil {
		if *data.Quantity <= 0 {
			return ErrInvalidQuantity
		}

		// Check stock availability
		var product productmodel.Product
		if err := biz.store.GetDB().Where("id = ?", orderItem.ProductId).First(&product).Error; err != nil {
			return ErrProductNotFound
		}

		if product.Stock < *data.Quantity {
			return ErrInsufficientStock
		}
	}

	// 4. Update order item
	return store.UpdateOrderItem(ctx, id, data)
}

func (biz *updateOrderItemBusiness) UpdateOrderItemQuantity(ctx context.Context, id int, quantity int) error {
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

	// Only allow updating items in pending orders
	if order.Status != ordermodel.OrderStatusPending {
		return ErrOrderNotPending
	}

	// 3. Validate quantity
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	// 4. Check stock availability
	var product productmodel.Product
	if err := biz.store.GetDB().Where("id = ?", orderItem.ProductId).First(&product).Error; err != nil {
		return ErrProductNotFound
	}

	if product.Stock < quantity {
		return ErrInsufficientStock
	}

	// 5. Update quantity and price
	return store.UpdateOrderItemQuantity(ctx, id, quantity)
}
