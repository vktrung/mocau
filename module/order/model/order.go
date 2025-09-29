package model

import (
	"mocau-backend/common"
	usermodel "mocau-backend/module/user/model"
	"time"
)

const EntityName = "Order"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"   // Chờ xử lý
	OrderStatusConfirmed OrderStatus = "confirmed" // Đã xác nhận
	OrderStatusCompleted OrderStatus = "completed" // Hoàn thành
	OrderStatusCancelled OrderStatus = "cancelled" // Đã hủy
)

func (status OrderStatus) String() string {
	return string(status)
}

type Order struct {
	common.SQLModel
	OrderNumber   string      `json:"order_number" gorm:"column:order_number;size:50;uniqueIndex;not null"`
	Status        OrderStatus `json:"status" gorm:"column:status;size:20;default:'pending'"`
	TotalAmount   float64     `json:"total_amount" gorm:"column:total_amount;type:decimal(12,2)"`
	
	// Thông tin khách hàng
	CustomerName    string `json:"customer_name" gorm:"column:customer_name;size:255;not null"`
	CustomerPhone   string `json:"customer_phone" gorm:"column:customer_phone;size:20;not null"`
	CustomerEmail   string `json:"customer_email" gorm:"column:customer_email;size:100"`
	ShippingAddress string `json:"shipping_address" gorm:"column:shipping_address;type:text;not null"`
	Notes           string `json:"notes" gorm:"column:notes;type:text"`
	
	// Thông tin xử lý
	ProcessedBy   *int       `json:"processed_by" gorm:"column:processed_by"` // ID nhân viên xử lý
	ProcessedAt   *time.Time `json:"processed_at" gorm:"column:processed_at"`
	CompletedAt   *time.Time `json:"completed_at" gorm:"column:completed_at"`
	
	// Relations
	Processor  *usermodel.User `json:"processor,omitempty" gorm:"foreignKey:ProcessedBy"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderCreate struct {
	common.SQLModel `json:",inline"`
	TotalAmount     float64 `json:"total_amount" gorm:"column:total_amount;type:decimal(12,2)"`
	
	// Thông tin khách hàng
	CustomerName    string `json:"customer_name" gorm:"column:customer_name;size:255;not null"`
	CustomerPhone   string `json:"customer_phone" gorm:"column:customer_phone;size:20;not null"`
	CustomerEmail   string `json:"customer_email" gorm:"column:customer_email;size:100"`
	ShippingAddress string `json:"shipping_address" gorm:"column:shipping_address;type:text;not null"`
	Notes           string `json:"notes" gorm:"column:notes;type:text"`
	
	// Danh sách sản phẩm (có thể tạo cùng lúc hoặc tạo riêng)
	OrderItems []OrderItemCreate `json:"order_items,omitempty" gorm:"-"`
}

type OrderItemCreate struct {
	ProductId int     `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Price     float64 `json:"price,omitempty"` // Optional, sẽ lấy giá hiện tại nếu không có
}

func (OrderCreate) TableName() string {
	return Order{}.TableName()
}

type OrderUpdate struct {
	Status          *OrderStatus `json:"status" gorm:"column:status;size:20"`
	CustomerName    *string      `json:"customer_name" gorm:"column:customer_name;size:255"`
	CustomerPhone   *string      `json:"customer_phone" gorm:"column:customer_phone;size:20"`
	CustomerEmail   *string      `json:"customer_email" gorm:"column:customer_email;size:100"`
	ShippingAddress *string      `json:"shipping_address" gorm:"column:shipping_address;type:text"`
	Notes           *string      `json:"notes" gorm:"column:notes;type:text"`
	ProcessedBy     *int         `json:"processed_by" gorm:"column:processed_by"`
	ProcessedAt     *time.Time   `json:"processed_at" gorm:"column:processed_at"`
	CompletedAt     *time.Time   `json:"completed_at" gorm:"column:completed_at"`
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}

type OrderFilter struct {
	Status      OrderStatus `json:"status" form:"status"`
	ProcessedBy int         `json:"processed_by" form:"processed_by"`
	CustomerPhone string    `json:"customer_phone" form:"customer_phone"`
	OrderNumber string      `json:"order_number" form:"order_number"`
}


