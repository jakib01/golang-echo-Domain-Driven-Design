package Usecase

import (
	"simple-rest-go-echo/Infra"
	"simple-rest-go-echo/Models"
)

type ProductUseCase struct {
	IProductRepository Models.IProductRepository
}

func NewProductUseCase() *ProductUseCase {
	return &ProductUseCase{
		IProductRepository: Infra.NewIProductRepository(),
	}
}

func (p ProductUseCase) Create(product *[]Models.Product) error {
	println(product)
	err := p.IProductRepository.CreateMySql(product)
	if err != nil {
		return err
	}
	return nil
}
