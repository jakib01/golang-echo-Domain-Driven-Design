package Routes

import (
	"gorm.io/gorm"
	"simple-rest-go-echo/Handler"

	"github.com/labstack/echo/v4"
)

// SetupRoutes setup for routing
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	productHandler := Handler.NewProductHandler(db)

	// Product route
	e.GET("/products", productHandler.GetAllProduct)
	e.POST("/product-create", productHandler.CreateProduct)
}
