package Models

import "time"

type Product struct {
	Id          int       `json:"product_id" gorm:"primaryKey;autoIncrement:true"`
	Name        string    `json:"product_name" gorm:"product_name"`
	Description string    `json:"description" gorm:"description"`
	Price       int       `json:"price" gorm:"price"`
	Stock       int       `json:"stock_quantity" gorm:"stock_quantity"`
	CategoryId  int       `json:"category_id" gorm:"category_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:time"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:time"`
}

type IProductUseCase interface {
	Create(product *[]Product) error
}

type IProductRepository interface {
	CreateMySql(product *[]Product) error
}
