package model

import "mocau-backend/common"

type Product struct {
    common.SQLModel
    Name        string   `json:"name" gorm:"column:name;size:255;not null"`
    Description string   `json:"description" gorm:"column:description;type:text"`
    Price       float64  `json:"price" gorm:"column:price;type:decimal(12,2)"`
    Stock       int      `json:"stock" gorm:"column:stock;default:0"`
    CategoryId  *int     `json:"category_id" gorm:"column:category_id"`
}

func (Product) TableName() string { return "products" }

type ProductCreate struct {
    common.SQLModel `json:",inline"`
    Name            string  `json:"name" gorm:"column:name;not null"`
    Description     string  `json:"description" gorm:"column:description"`
    Price           float64 `json:"price" gorm:"column:price"`
    Stock           int     `json:"stock" gorm:"column:stock"`
    CategoryId      *int    `json:"category_id" gorm:"column:category_id"`
}

func (ProductCreate) TableName() string { return Product{}.TableName() }

type ProductUpdate struct {
    Name        *string  `json:"name" gorm:"column:name"`
    Description *string  `json:"description" gorm:"column:description"`
    Price       *float64 `json:"price" gorm:"column:price"`
    Stock       *int     `json:"stock" gorm:"column:stock"`
    CategoryId  *int     `json:"category_id" gorm:"column:category_id"`
}

func (ProductUpdate) TableName() string { return Product{}.TableName() }


