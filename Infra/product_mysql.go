package Infra

import (
	"gorm.io/gorm"
	"simple-rest-go-echo/Models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewIProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p ProductRepository) TxStart() (*gorm.DB, error) {
	tx := p.db.Begin()
	return tx, tx.Error
}

func (p ProductRepository) TxCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (p ProductRepository) TxRollback(tx *gorm.DB) {
	tx.Rollback()
}

func (p ProductRepository) FetchAllProduct() ([]Models.Product, error) {
	var products []Models.Product
	result := p.db.Find(&products)
	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}

func (p ProductRepository) FetchProductById(productId int64) (Models.Product, error) {
	result := Models.Product{}
	err := p.db.
		Table("products").
		Where("product_id = ?", productId).
		First(&result).Error
	return result, err
}

func (p ProductRepository) CreateProduct(request *Models.Product) error {
	return p.db.Create(request).Error
}
