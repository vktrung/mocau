package biz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/model"
	"mocau-backend/module/orderitem/storage"
	ordermodel "mocau-backend/module/order/model"
	productmodel "mocau-backend/module/product/model"
)

type CreateOrderItemStorage interface {
	GetDB() *gorm.DB
}

type createOrderItemBusiness struct {
	store CreateOrderItemStorage
}

func NewCreateOrderItemBusiness(store CreateOrderItemStorage) *createOrderItemBusiness {
	return &createOrderItemBusiness{store: store}
}

func (biz *createOrderItemBusiness) CreateOrderItem(ctx context.Context, data *model.OrderItemCreate) error {
	// Business Rules Validation

	// 1. Validate required fields
	if data.OrderId <= 0 {
		return ErrInvalidOrderId
	}
	if data.ProductId <= 0 {
		return ErrInvalidProductId
	}
	if data.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	// 2. Check if order exists and is in pending status
	var order ordermodel.Order
	if err := biz.store.GetDB().Where("id = ?", data.OrderId).First(&order).Error; err != nil {
		return ErrOrderNotFound
	}

	// Only allow adding items to pending orders
	if order.Status != ordermodel.OrderStatusPending {
		return ErrOrderNotPending
	}

	// 3. Check if product exists and has stock
	var product productmodel.Product
	if err := biz.store.GetDB().Where("id = ?", data.ProductId).First(&product).Error; err != nil {
		return ErrProductNotFound
	}

	// Check stock availability
	if product.Stock < data.Quantity {
		return ErrInsufficientStock
	}

	// 4. Create order item through storage
	store := storage.NewSQLStore(biz.store.GetDB())
	if err := store.CreateOrderItem(ctx, data); err != nil {
		return err
	}

	// 5. Deduct stock from product
	if err := biz.store.GetDB().Model(&productmodel.Product{}).
		Where("id = ?", data.ProductId).
		Update("stock", product.Stock-data.Quantity).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

// Business Rule Errors
var (
	ErrInvalidOrderId    = common.NewCustomError(errors.New("invalid order id"), "invalid order id", "ErrInvalidOrderId")
	ErrOrderNotFound     = common.NewCustomError(errors.New("order not found"), "order not found", "ErrOrderNotFound")
	ErrOrderNotPending   = common.NewCustomError(errors.New("order is not pending"), "can only add items to pending orders", "ErrOrderNotPending")
	ErrInvalidProductId  = common.NewCustomError(errors.New("invalid product id"), "invalid product id", "ErrInvalidProductId")
	ErrInvalidQuantity   = common.NewCustomError(errors.New("invalid quantity"), "invalid quantity", "ErrInvalidQuantity")
	ErrProductNotFound   = common.NewCustomError(errors.New("product not found"), "product not found", "ErrProductNotFound")
	ErrInsufficientStock = common.NewCustomError(errors.New("insufficient stock"), "insufficient stock", "ErrInsufficientStock")
)
