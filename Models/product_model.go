package Models

import "golang-echo-Domain-Driven-Design/Common"

type Product struct {
	ProductTable
}

type ProductTable struct {
	ProductId     int     `json:"product_id" gorm:"primaryKey;autoIncrement:true"`
	ProductName   string  `json:"product_name" gorm:"product_name"`
	Description   string  `json:"description" gorm:"description"`
	Price         float64 `json:"price" gorm:"price"`
	StockQuantity int     `json:"stock_quantity" gorm:"stock_quantity"`
	CategoryId    int     `json:"category_id" gorm:"category_id"`
	Common.Times
}

type SingleProductInput struct {
	ProductId int64 `json:"product_id" param:"ProductId" validate:"required"`
}

type IProductUseCase interface {
	Create(request *Product) error
	GetAllProduct() ([]Product, error)
	GetProduct(request *SingleProductInput) (Product, error)
	Update(productId string, request *Product) error
	Delete(productId string) error
}

type IProductRepository interface {
	Common.Repository
	CreateProduct(request *Product) error
	FetchAllProduct() ([]Product, error)
	FetchProductById(productId int64) (Product, error)
	UpdateProduct(productId int64, request *Product) error
	DeleteProduct(productId int64) error
}
