package Routes

import (
	"simple-rest-go-echo/Handler"

	"github.com/labstack/echo/v4"
)

// SetupRoutes setup for routing
func SetupRoutes(e *echo.Echo) {
	productHandler := Handler.NewProductHandler()

	// Product route
	e.GET("/product", productHandler.GetProduct)
	e.POST("/product-create", productHandler.CreateProduct)
}
