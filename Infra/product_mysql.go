package Infra

import "simple-rest-go-echo/Models"

type ProductRepository struct{}

func NewIProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (p ProductRepository) CreateMySql(product *[]Models.Product) error {
	println(product)
	return nil
}
