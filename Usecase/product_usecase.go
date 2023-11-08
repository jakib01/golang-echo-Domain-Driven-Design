package Usecase

import (
	"simple-rest-go-echo/Models"
)

type ProductUseCase11 struct {
	IProductRepo Models.IRepo
}

func NewProductUseCase() Models.IProductUseCase {
	return &ProductUseCase11{}
}

func (p ProductUseCase11) Create(product *[]Models.Product) error {

	//err := p.IProductRepo.CreateMySql(product)
	//if err != nil {
	//	return err
	//}
	return nil
}
