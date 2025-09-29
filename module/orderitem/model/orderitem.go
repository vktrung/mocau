package model

import (
	"mocau-backend/common"
	productmodel "mocau-backend/module/product/model"
	ordermodel "mocau-backend/module/order/model"
)

const EntityName = "OrderItem"

type OrderItem struct {
	common.SQLModel
	OrderId   int     `json:"order_id" gorm:"column:order_id;not null"`
	ProductId int     `json:"product_id" gorm:"column:product_id;not null"`
	Quantity  int     `json:"quantity" gorm:"column:quantity;not null"`
	Price     float64 `json:"price" gorm:"column:price;type:decimal(12,2);not null"`
	
	// Relations
	Order   *ordermodel.Order                `json:"order,omitempty" gorm:"foreignKey:OrderId"`
	Product *productmodel.Product            `json:"product,omitempty" gorm:"foreignKey:ProductId"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

type OrderItemCreate struct {
	common.SQLModel `json:",inline"`
	OrderId         int     `json:"order_id" gorm:"column:order_id;not null"`
	ProductId       int     `json:"product_id" gorm:"column:product_id;not null"`
	Quantity        int     `json:"quantity" gorm:"column:quantity;not null"`
	Price           float64 `json:"price" gorm:"column:price;type:decimal(12,2);not null"`
}

func (OrderItemCreate) TableName() string {
	return OrderItem{}.TableName()
}

type OrderItemUpdate struct {
	Quantity *int     `json:"quantity" gorm:"column:quantity"`
	Price    *float64 `json:"price" gorm:"column:price;type:decimal(12,2)"`
}

func (OrderItemUpdate) TableName() string {
	return OrderItem{}.TableName()
}
