package csv

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateClientsBatch(clients []*Client) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateClientsBatch(clients []*Client) error {
	if len(clients) == 0 {
		return nil
	}
	return r.db.CreateInBatches(clients, len(clients)).Error
}
