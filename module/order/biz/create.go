package biz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
	"mocau-backend/module/order/storage"
	productmodel "mocau-backend/module/product/model"
)

type CreateOrderStorage interface {
	GetDB() *gorm.DB
}

type createOrderBusiness struct {
	store CreateOrderStorage
}

func NewCreateOrderBusiness(store CreateOrderStorage) *createOrderBusiness {
	return &createOrderBusiness{store: store}
}

func (biz *createOrderBusiness) CreateOrder(ctx context.Context, data *model.OrderCreate) error {
	// Business Rules Validation
	
	// 1. Validate required fields
	if data.CustomerName == "" {
		return ErrCustomerNameRequired
	}
	if data.CustomerPhone == "" {
		return ErrCustomerPhoneRequired
	}
	if data.ShippingAddress == "" {
		return ErrShippingAddressRequired
	}

	// 2. Validate and calculate total amount from order items
	totalAmount := 0.0
	if len(data.OrderItems) > 0 {
		for i, item := range data.OrderItems {
			if item.ProductId <= 0 {
				return ErrInvalidProductId
			}
			if item.Quantity <= 0 {
				return ErrInvalidQuantity
			}

			// Check product exists and get current price
			var product productmodel.Product
			if err := biz.store.GetDB().Where("id = ?", item.ProductId).First(&product).Error; err != nil {
				return ErrProductNotFound
			}

			// Check stock availability
			if product.Stock < item.Quantity {
				return ErrInsufficientStock
			}

			// Use current product price
			data.OrderItems[i].Price = product.Price
			totalAmount += product.Price * float64(item.Quantity)
		}
	}

	// 3. Set calculated total amount
	data.TotalAmount = totalAmount

	// 4. Create order through storage
	store := storage.NewSQLStore(biz.store.GetDB())
	return store.CreateOrder(ctx, data)
}

// Business Rule Errors
var (
	ErrCustomerNameRequired    = common.NewCustomError(errors.New("customer name is required"), "customer name is required", "ErrCustomerNameRequired")
	ErrCustomerPhoneRequired   = common.NewCustomError(errors.New("customer phone is required"), "customer phone is required", "ErrCustomerPhoneRequired")
	ErrShippingAddressRequired = common.NewCustomError(errors.New("shipping address is required"), "shipping address is required", "ErrShippingAddressRequired")
	ErrOrderItemsRequired      = common.NewCustomError(errors.New("order items are required"), "order items are required", "ErrOrderItemsRequired")
	ErrInvalidProductId        = common.NewCustomError(errors.New("invalid product id"), "invalid product id", "ErrInvalidProductId")
	ErrInvalidQuantity         = common.NewCustomError(errors.New("invalid quantity"), "invalid quantity", "ErrInvalidQuantity")
	ErrInvalidPrice            = common.NewCustomError(errors.New("invalid price"), "invalid price", "ErrInvalidPrice")
	ErrProductNotFound         = common.NewCustomError(errors.New("product not found"), "product not found", "ErrProductNotFound")
	ErrInsufficientStock       = common.NewCustomError(errors.New("insufficient stock"), "insufficient stock", "ErrInsufficientStock")
	ErrInvalidTotalAmount      = common.NewCustomError(errors.New("invalid total amount"), "invalid total amount", "ErrInvalidTotalAmount")
)
