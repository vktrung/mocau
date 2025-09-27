package storage

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/module/blog/model"
)

type Store interface {
	CreateBlog(ctx context.Context, data *model.BlogCreate) error
	GetBlog(ctx context.Context, conditions map[string]interface{}) (*model.Blog, error)
	GetBlogWithAuthor(ctx context.Context, id int) (*model.BlogWithAuthor, error)
	ListBlogs(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) ([]model.Blog, error)
	UpdateBlog(ctx context.Context, id int, data *model.BlogUpdate) error
	DeleteBlog(ctx context.Context, id int) error
}

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
