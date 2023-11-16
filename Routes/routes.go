package Routes

import (
	"golang-echo-Domain-Driven-Design/Handler"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

// SetupRoutes setup for routing
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	productHandler := Handler.NewProductHandler(db)

	// Product route
	e.GET("/products", productHandler.GetAllProduct)
	e.GET("/product/:ProductId", productHandler.GetProduct)
	e.POST("/product-create", productHandler.CreateProduct)
	e.PUT("/product-update/:productId", productHandler.UpdateProduct)
	e.DELETE("/product-delete/:productId", productHandler.DeleteProduct)
}
