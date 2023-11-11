package Models

import "simple-rest-go-echo/Common"

type Product struct {
	ProductTable
}

type ProductTable struct {
	Id            int    `json:"product_id" gorm:"primaryKey;autoIncrement:true"`
	ProductName   string `json:"product_name" gorm:"product_name"`
	Description   string `json:"description" gorm:"description"`
	Price         int    `json:"price" gorm:"price"`
	StockQuantity int    `json:"stock_quantity" gorm:"stock_quantity"`
	CategoryId    int    `json:"category_id" gorm:"category_id"`
	Common.Times
}

type IProductUseCase interface {
	Create(request *Product) error
}

type IProductRepository interface {
	Common.Repository
	CreateProduct(request *Product) error
}
