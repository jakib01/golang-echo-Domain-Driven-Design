package Infra

import (
	"gorm.io/gorm"
	"simple-rest-go-echo/Models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewIProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p ProductRepository) TxStart() (*gorm.DB, error) {
	tx := p.db.Begin()
	return tx, tx.Error
}

func (p ProductRepository) TxCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (p ProductRepository) TxRollback(tx *gorm.DB) {
	tx.Rollback()
}

func (p ProductRepository) CreateProduct(request *Models.Product) error {
	return p.db.Create(request).Error
}
