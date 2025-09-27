package model

import "mocau-backend/common"

type Blog struct {
	common.SQLModel
	Title    string        `json:"title" gorm:"column:title;size:255;not null"`
	Content  string        `json:"content" gorm:"column:content;type:longtext"` // HTML content
	AuthorId int           `json:"author_id" gorm:"column:author_id;not null"`
	Image    *common.Image `json:"image" gorm:"column:image;type:json"`
	Status   string        `json:"status" gorm:"column:status;default:'draft'"` // draft, published
}

func (Blog) TableName() string { return "blogs" }

type BlogCreate struct {
	common.SQLModel `json:",inline"`
	Title           string        `json:"title" gorm:"column:title;not null"`
	Content         string        `json:"content" gorm:"column:content"`
	AuthorId        int           `json:"author_id" gorm:"column:author_id;not null"`
	Image           *common.Image `json:"image" gorm:"column:image;type:json"`
	Status          string        `json:"status" gorm:"column:status;default:'draft'"`
}

func (BlogCreate) TableName() string { return Blog{}.TableName() }

type BlogUpdate struct {
	Title   *string       `json:"title" gorm:"column:title"`
	Content *string       `json:"content" gorm:"column:content"`
	Image   *common.Image `json:"image" gorm:"column:image;type:json"`
	Status  *string       `json:"status" gorm:"column:status"`
}

func (BlogUpdate) TableName() string { return Blog{}.TableName() }

// BlogWithAuthor for response with author info
type BlogWithAuthor struct {
	Blog
	Author struct {
		Id       int    `json:"id"`
		FullName string `json:"full_name"`
	} `json:"author"`
}
