package Usecase

import (
	"gorm.io/gorm"
	"simple-rest-go-echo/Common"
	"simple-rest-go-echo/Infra"
	"simple-rest-go-echo/Models"
	"time"
)

type ProductUseCase struct {
	IProductRepository Models.IProductRepository
}

func NewProductUseCase(db *gorm.DB) *ProductUseCase {
	return &ProductUseCase{
		IProductRepository: Infra.NewIProductRepository(db),
	}
}

func (p ProductUseCase) Create(request *Models.Product) error {
	// prepare product data
	product := &Models.Product{
		ProductTable: Models.ProductTable{
			ProductName:   request.ProductName,
			Description:   request.Description,
			Price:         request.Price,
			StockQuantity: request.StockQuantity,
			CategoryId:    request.CategoryId,
			Times:         Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := p.IProductRepository.CreateProduct(product); err != nil {
		return err
	}
	return nil
}
