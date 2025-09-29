package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/orderitem/model"
	productmodel "mocau-backend/module/product/model"
)

func (s *sqlStore) UpdateOrderItem(ctx context.Context, id int, data *model.OrderItemUpdate) error {
	updates := make(map[string]interface{})

	if data.Quantity != nil {
		updates["quantity"] = *data.Quantity
	}
	if data.Price != nil {
		updates["price"] = *data.Price
	}

	if err := s.db.Model(&model.OrderItem{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateOrderItemQuantity(ctx context.Context, id int, quantity int) error {
	// Get current order item
	var orderItem model.OrderItem
	if err := s.db.Where("id = ?", id).First(&orderItem).Error; err != nil {
		return common.RecordNotFound
	}

	// Get current product price
	var product productmodel.Product
	if err := s.db.Where("id = ?", orderItem.ProductId).First(&product).Error; err != nil {
		return common.RecordNotFound
	}

	// Update quantity and price
	updates := map[string]interface{}{
		"quantity": quantity,
		"price":    product.Price, // Always use current product price
	}

	if err := s.db.Model(&model.OrderItem{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
