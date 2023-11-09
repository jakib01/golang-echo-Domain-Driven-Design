package Infra

import (
	"simple-rest-go-echo/Config"
	"simple-rest-go-echo/Models"
)

type ProductRepository struct{}

func NewIProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (p ProductRepository) CreateProduct(request *Models.Product) error {
	db := Config.GetDB()
	return db.Create(request).Error
}
