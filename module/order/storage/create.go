package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
	orderitemmodel "mocau-backend/module/orderitem/model"
	"time"
)

func (s *sqlStore) CreateOrder(ctx context.Context, data *model.OrderCreate) error {
	db := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	// Generate order number
	orderNumber := s.generateOrderNumber()
	
	// Create order
	order := &model.Order{
		OrderNumber:     orderNumber,
		Status:          model.OrderStatusPending,
		TotalAmount:     data.TotalAmount,
		CustomerName:    data.CustomerName,
		CustomerPhone:   data.CustomerPhone,
		CustomerEmail:   data.CustomerEmail,
		ShippingAddress: data.ShippingAddress,
		Notes:           data.Notes,
	}

	if err := db.Create(order).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

		// Create order items if provided
		if len(data.OrderItems) > 0 {
			for _, item := range data.OrderItems {
				orderItem := &orderitemmodel.OrderItem{
					OrderId:   order.Id,
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
					Price:     item.Price,
				}

				if err := db.Create(orderItem).Error; err != nil {
					db.Rollback()
					return common.ErrDB(err)
				}
			}
		}

	return db.Commit().Error
}

func (s *sqlStore) generateOrderNumber() string {
	// Generate order number: ORD + timestamp + random
	timestamp := time.Now().Format("20060102150405")
	return "ORD" + timestamp
}
