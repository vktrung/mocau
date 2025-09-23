package model

import (
    "errors"
    "mocau-backend/common"
)

var (
    ErrCategoryNameCannotBeEmpty = errors.New("category name cannot be empty")
)

const (
	EntityName = "Category"
)

type Category struct {
	common.SQLModel
	CategoryName string `json:"category_name" gorm:"column:category_name;"`
	Description  string `json:"description" gorm:"column:description"`
	Status       string `json:"status" gorm:"column:status;"`
}

func (Category) TableName() string { return "Categories" }

type CategoryCreate struct {
    common.SQLModel `json:",inline"`
    CategoryName    string `json:"category_name" gorm:"column:category_name;"`
    Description     string `json:"description" gorm:"column:description"`
    Status          string `json:"status" gorm:"column:status;"`
}

func (CategoryCreate) TableName() string { return Category{}.TableName() }

type CategoryUpdate struct {
    CategoryName *string `json:"category_name" gorm:"column:category_name"`
    Description  *string `json:"description" gorm:"column:description"`
    Status       *string `json:"status" gorm:"column:status"`
}

func (CategoryUpdate) TableName() string { return Category{}.TableName() }
