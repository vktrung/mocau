package storage

import (
	"context"
	"gorm.io/gorm"
	"mocau-backend/module/order/model"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (s *sqlStore) GetDB() *gorm.DB {
	return s.db
}
