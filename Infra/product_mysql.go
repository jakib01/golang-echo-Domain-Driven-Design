package Infra

import (
	"simple-rest-go-echo/Models"
)

type ProductUseCase22 struct {
	IProductRepo Models.IRepo
}

func NewProductRepo() Models.IRepo {
	return &ProductUseCase22{}
}

func (p ProductUseCase22) CreateMySql(product *[]Models.Product) error {

	//fmt.Println(product)

	return nil
}
