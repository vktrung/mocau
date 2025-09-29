package biz

import (
	"context"
	"errors"
	"mocau-backend/common"
	"mocau-backend/module/order/model"
	"mocau-backend/module/order/storage"
	usermodel "mocau-backend/module/user/model"
)

type UpdateOrderStorage interface {
	GetDB() *gorm.DB
}

type updateOrderBusiness struct {
	store UpdateOrderStorage
}

func NewUpdateOrderBusiness(store UpdateOrderStorage) *updateOrderBusiness {
	return &updateOrderBusiness{store: store}
}

func (biz *updateOrderBusiness) UpdateOrder(ctx context.Context, id int, data *model.OrderUpdate) error {
	store := storage.NewSQLStore(biz.store.GetDB())

	// Business Rules Validation

	// 1. Check if order exists
	order, err := store.GetOrder(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	// 2. Validate status transitions
	if data.Status != nil {
		if !biz.isValidStatusTransition(order.Status, *data.Status) {
			return ErrInvalidStatusTransition
		}
	}

	// 3. Validate processed_by (must be admin user)
	if data.ProcessedBy != nil {
		var user usermodel.User
		if err := biz.store.GetDB().Where("id = ? AND role = ?", *data.ProcessedBy, usermodel.RoleAdmin.String()).First(&user).Error; err != nil {
			return ErrInvalidProcessor
		}
	}

	// 4. Validate customer phone format (basic validation)
	if data.CustomerPhone != nil {
		if len(*data.CustomerPhone) < 10 {
			return ErrInvalidPhoneFormat
		}
	}

	return store.UpdateOrder(ctx, id, data)
}

func (biz *updateOrderBusiness) UpdateOrderStatus(ctx context.Context, id int, status model.OrderStatus, processedBy *int) error {
	store := storage.NewSQLStore(biz.store.GetDB())

	// Business Rules Validation

	// 1. Check if order exists
	order, err := store.GetOrder(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	// 2. Validate status transitions
	if !biz.isValidStatusTransition(order.Status, status) {
		return ErrInvalidStatusTransition
	}

	// 3. Validate processed_by (must be admin user)
	if processedBy != nil {
		var user usermodel.User
		if err := biz.store.GetDB().Where("id = ? AND role = ?", *processedBy, usermodel.RoleAdmin.String()).First(&user).Error; err != nil {
			return ErrInvalidProcessor
		}
	}

	return store.UpdateOrderStatus(ctx, id, status, processedBy)
}

// Business Rules for Status Transitions
func (biz *updateOrderBusiness) isValidStatusTransition(currentStatus, newStatus model.OrderStatus) bool {
	// Define valid status transitions
	validTransitions := map[model.OrderStatus][]model.OrderStatus{
		model.OrderStatusPending:   {model.OrderStatusConfirmed, model.OrderStatusCancelled},
		model.OrderStatusConfirmed: {model.OrderStatusCompleted, model.OrderStatusCancelled},
		model.OrderStatusCompleted: {}, // No transitions from completed
		model.OrderStatusCancelled: {}, // No transitions from cancelled
	}

	allowedStatuses, exists := validTransitions[currentStatus]
	if !exists {
		return false
	}

	for _, allowedStatus := range allowedStatuses {
		if allowedStatus == newStatus {
			return true
		}
	}

	return false
}

// Business Rule Errors
var (
	ErrInvalidStatusTransition = common.NewCustomError(errors.New("invalid status transition"), "invalid status transition", "ErrInvalidStatusTransition")
	ErrInvalidProcessor        = common.NewCustomError(errors.New("invalid processor"), "processor must be an admin user", "ErrInvalidProcessor")
	ErrInvalidPhoneFormat      = common.NewCustomError(errors.New("invalid phone format"), "phone number must be at least 10 digits", "ErrInvalidPhoneFormat")
)
