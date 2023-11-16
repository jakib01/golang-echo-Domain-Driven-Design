package Usecase

import (
	"golang-echo-Domain-Driven-Design/Common"
	"golang-echo-Domain-Driven-Design/Infra"
	"golang-echo-Domain-Driven-Design/Models"
	"gorm.io/gorm"
	"strconv"
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

func (p ProductUseCase) GetAllProduct() ([]Models.Product, error) {
	product, err := p.IProductRepository.FetchAllProduct()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p ProductUseCase) GetProduct(request *Models.SingleProductInput) (Models.Product, error) {
	product, err := p.IProductRepository.FetchProductById(request.ProductId)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p ProductUseCase) Create(request *Models.Product) error {
	tx, txErr := p.IProductRepository.TxStart()
	if txErr != nil {
		return txErr
	}

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
		p.IProductRepository.TxRollback(tx)
		return err
	}

	if err := p.IProductRepository.TxCommit(tx); err != nil {
		p.IProductRepository.TxRollback(tx)
		return err
	}
	return nil
}

func (p ProductUseCase) Update(productId string, request *Models.Product) error {
	tx, txErr := p.IProductRepository.TxStart()
	if txErr != nil {
		return txErr
	}
	productID, _ := strconv.Atoi(productId)

	product, err := p.IProductRepository.FetchProductById(int64(productID))
	if err != nil {
		return err
	}

	// prepare product data for update
	productUpdate := &Models.Product{
		ProductTable: Models.ProductTable{
			ProductName:   request.ProductName,
			Description:   request.Description,
			Price:         request.Price,
			StockQuantity: request.StockQuantity,
			CategoryId:    request.CategoryId,
			Times:         Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := p.IProductRepository.UpdateProduct(int64(product.ProductId), productUpdate); err != nil {
		p.IProductRepository.TxRollback(tx)
		return err
	}

	if err := p.IProductRepository.TxCommit(tx); err != nil {
		p.IProductRepository.TxRollback(tx)
		return err
	}
	return nil
}

func (p ProductUseCase) Delete(productId string) error {
	tx, txErr := p.IProductRepository.TxStart()
	if txErr != nil {
		return txErr
	}
	productID, _ := strconv.Atoi(productId)

	if err := p.IProductRepository.DeleteProduct(int64(productID)); err != nil {
		p.IProductRepository.TxRollback(tx)
		return err
	}

	if err := p.IProductRepository.TxCommit(tx); err != nil {
		p.IProductRepository.TxRollback(tx)
		return err
	}

	return nil
}
