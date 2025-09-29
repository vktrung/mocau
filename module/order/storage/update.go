package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
	"time"
)

func (s *sqlStore) UpdateOrder(ctx context.Context, id int, data *model.OrderUpdate) error {
	updates := make(map[string]interface{})

	if data.Status != nil {
		updates["status"] = *data.Status
		
		// Set timestamps based on status
		now := time.Now()
		switch *data.Status {
		case model.OrderStatusConfirmed:
			if data.ProcessedAt == nil {
				updates["processed_at"] = now
			}
		case model.OrderStatusCompleted:
			if data.CompletedAt == nil {
				updates["completed_at"] = now
			}
		}
	}

	if data.CustomerName != nil {
		updates["customer_name"] = *data.CustomerName
	}
	if data.CustomerPhone != nil {
		updates["customer_phone"] = *data.CustomerPhone
	}
	if data.CustomerEmail != nil {
		updates["customer_email"] = *data.CustomerEmail
	}
	if data.ShippingAddress != nil {
		updates["shipping_address"] = *data.ShippingAddress
	}
	if data.Notes != nil {
		updates["notes"] = *data.Notes
	}
	if data.ProcessedBy != nil {
		updates["processed_by"] = *data.ProcessedBy
	}
	if data.ProcessedAt != nil {
		updates["processed_at"] = *data.ProcessedAt
	}
	if data.CompletedAt != nil {
		updates["completed_at"] = *data.CompletedAt
	}

	if err := s.db.Model(&model.Order{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateOrderStatus(ctx context.Context, id int, status model.OrderStatus, processedBy *int) error {
	updates := map[string]interface{}{
		"status": status,
	}

	now := time.Now()
	switch status {
	case model.OrderStatusConfirmed:
		updates["processed_at"] = now
		if processedBy != nil {
			updates["processed_by"] = *processedBy
		}
	case model.OrderStatusCompleted:
		updates["completed_at"] = now
		if processedBy != nil {
			updates["processed_by"] = *processedBy
		}
	}

	if err := s.db.Model(&model.Order{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
