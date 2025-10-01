package biz

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/common"
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

	// 3. Validate quantity if provided and handle stock
	if data.Quantity != nil {
		if *data.Quantity <= 0 {
			return ErrInvalidQuantity
		}

		// Check stock availability and calculate stock difference
		var product productmodel.Product
		if err := biz.store.GetDB().Where("id = ?", orderItem.ProductId).First(&product).Error; err != nil {
			return ErrProductNotFound
		}

		// Calculate stock difference (new quantity - old quantity)
		stockDifference := *data.Quantity - orderItem.Quantity
		
		// Check if we have enough stock for the difference
		if product.Stock < stockDifference {
			return ErrInsufficientStock
		}

		// Update order item first
		if err := store.UpdateOrderItem(ctx, id, data); err != nil {
			return err
		}

		// Update stock based on the difference
		if stockDifference != 0 {
			if err := biz.store.GetDB().Model(&productmodel.Product{}).
				Where("id = ?", orderItem.ProductId).
				Update("stock", product.Stock-stockDifference).Error; err != nil {
				return common.ErrDB(err)
			}
		}

		return nil
	}

	// 4. Update order item (no quantity change)
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

	// 4. Check stock availability and calculate stock difference
	var product productmodel.Product
	if err := biz.store.GetDB().Where("id = ?", orderItem.ProductId).First(&product).Error; err != nil {
		return ErrProductNotFound
	}

	// Calculate stock difference (new quantity - old quantity)
	stockDifference := quantity - orderItem.Quantity
	
	// Check if we have enough stock for the difference
	if product.Stock < stockDifference {
		return ErrInsufficientStock
	}

	// 5. Update quantity and price
	if err := store.UpdateOrderItemQuantity(ctx, id, quantity); err != nil {
		return err
	}

	// 6. Update stock based on the difference
	if stockDifference != 0 {
		if err := biz.store.GetDB().Model(&productmodel.Product{}).
			Where("id = ?", orderItem.ProductId).
			Update("stock", product.Stock-stockDifference).Error; err != nil {
			return common.ErrDB(err)
		}
	}

	return nil
}
