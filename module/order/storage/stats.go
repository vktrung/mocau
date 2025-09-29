package storage

import (
	"context"
	"mocau-backend/module/order/model"
)

type OrderStats struct {
	TotalOrders     int     `json:"total_orders"`
	PendingOrders   int     `json:"pending_orders"`
	ConfirmedOrders int     `json:"confirmed_orders"`
	CompletedOrders int     `json:"completed_orders"`
	CancelledOrders int     `json:"cancelled_orders"`
	TotalRevenue    float64 `json:"total_revenue"`
	TodayOrders     int     `json:"today_orders"`
	TodayRevenue    float64 `json:"today_revenue"`
}

func (s *sqlStore) GetOrderStats(ctx context.Context) (*OrderStats, error) {
	var stats OrderStats

	// Count orders by status
	s.db.Model(&model.Order{}).Where("status = ?", model.OrderStatusPending).Count(&stats.PendingOrders)
	s.db.Model(&model.Order{}).Where("status = ?", model.OrderStatusConfirmed).Count(&stats.ConfirmedOrders)
	s.db.Model(&model.Order{}).Where("status = ?", model.OrderStatusCompleted).Count(&stats.CompletedOrders)
	s.db.Model(&model.Order{}).Where("status = ?", model.OrderStatusCancelled).Count(&stats.CancelledOrders)

	// Total orders
	stats.TotalOrders = stats.PendingOrders + stats.ConfirmedOrders + stats.CompletedOrders + stats.CancelledOrders

	// Total revenue (only from completed orders)
	s.db.Model(&model.Order{}).Where("status = ?", model.OrderStatusCompleted).Select("COALESCE(SUM(total_amount), 0)").Scan(&stats.TotalRevenue)

	// Today's orders and revenue
	s.db.Model(&model.Order{}).Where("DATE(created_at) = CURDATE()").Count(&stats.TodayOrders)
	s.db.Model(&model.Order{}).Where("DATE(created_at) = CURDATE() AND status = ?", model.OrderStatusCompleted).Select("COALESCE(SUM(total_amount), 0)").Scan(&stats.TodayRevenue)

	return &stats, nil
}
