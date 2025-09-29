package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/model"
	ordermodel "mocau-backend/module/order/model"
	productmodel "mocau-backend/module/product/model"
)

func (s *sqlStore) CreateOrderItem(ctx context.Context, data *model.OrderItemCreate) error {
	// Check if order exists
	var order ordermodel.Order
	if err := s.db.Where("id = ?", data.OrderId).First(&order).Error; err != nil {
		return common.ErrRecordNotFound
	}

	// Check if product exists and get current price
	var product productmodel.Product
	if err := s.db.Where("id = ?", data.ProductId).First(&product).Error; err != nil {
		return common.ErrRecordNotFound
	}

	// Check if order item already exists for this order and product
	var existingItem model.OrderItem
	if err := s.db.Where("order_id = ? AND product_id = ?", data.OrderId, data.ProductId).First(&existingItem).Error; err == nil {
		// Update existing item quantity
		existingItem.Quantity += data.Quantity
		return s.db.Save(&existingItem).Error
	}

	// Create new order item with current product price
	orderItem := &model.OrderItem{
		OrderId:   data.OrderId,
		ProductId: data.ProductId,
		Quantity:  data.Quantity,
		Price:     product.Price, // Use current product price
	}

	return s.db.Create(orderItem).Error
}
