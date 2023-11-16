package Infra

import (
	"golang-echo-Domain-Driven-Design/Models"
	"gorm.io/gorm"
	"time"
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

func (p ProductRepository) UpdateProduct(productId int64, request *Models.Product) error {
	return p.db.Model(&Models.Product{}).
		Where("product_id = ?", productId).
		Updates(map[string]interface{}{
			"product_name":   request.ProductName,
			"description":    request.Description,
			"price":          request.Price,
			"stock_quantity": request.StockQuantity,
			"category_id":    request.CategoryId,
			"updated_at":     time.Now(),
		}).Error
}

func (p ProductRepository) DeleteProduct(productId int64) error {
	return p.db.Delete(&Models.Product{}, "product_id = ?", productId).Error
}
