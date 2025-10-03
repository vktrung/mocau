package storage

import (
	"context"
	"mocau-backend/common"
	"mocau-backend/module/product/model"
	ordermodel "mocau-backend/module/order/model"
	"time"
)

func (s *sqlStore) GetRevenueGrowth(ctx context.Context) (*model.RevenueGrowth, error) {
	now := time.Now()
	
	// Tháng hiện tại
	currentStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	currentEnd := currentStart.AddDate(0, 1, 0).Add(-time.Second)
	
	// Tháng trước
	previousStart := currentStart.AddDate(0, -1, 0)
	previousEnd := currentStart.Add(-time.Second)

	var result model.RevenueGrowth

	// Tính doanh thu tháng hiện tại
	err := s.db.Model(&ordermodel.Order{}).
		Where("status = ? AND created_at >= ? AND created_at <= ?", 
			ordermodel.OrderStatusCompleted, currentStart, currentEnd).
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&result.CurrentMonth).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	// Tính doanh thu tháng trước
	err = s.db.Model(&ordermodel.Order{}).
		Where("status = ? AND created_at >= ? AND created_at <= ?", 
			ordermodel.OrderStatusCompleted, previousStart, previousEnd).
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&result.PreviousMonth).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	// Tính toán tăng trưởng
	result.GrowthAmount = result.CurrentMonth - result.PreviousMonth
	
	// Tính tỷ lệ tăng trưởng
	if result.PreviousMonth > 0 {
		result.GrowthRate = (result.GrowthAmount / result.PreviousMonth) * 100
	} else if result.CurrentMonth > 0 {
		// Nếu tháng trước = 0 và tháng này > 0, tăng trưởng 100%
		result.GrowthRate = 100.0
	} else {
		// Cả hai tháng đều = 0
		result.GrowthRate = 0.0
	}

	// Xác định hướng tăng trưởng
	if result.GrowthAmount > 0 {
		result.GrowthDirection = "up"
	} else if result.GrowthAmount < 0 {
		result.GrowthDirection = "down"
	} else {
		result.GrowthDirection = "stable"
	}

	return &result, nil
}
