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

func (p ProductRepository) CreateProduct(request *Models.Product) error {
	//db := Config.GetDB()
	return p.db.Create(request).Error
}
