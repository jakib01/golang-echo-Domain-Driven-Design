package Routes

import (
	"simple-rest-go-echo/Handler"

	"github.com/labstack/echo/v4"
)

// SetupRoutes setup for routing
func SetupRoutes(e *echo.Echo) {
	productHandler := Handler.NewProductHandler()

	e.GET("/course", Handler.GetCourse)
	e.POST("/course", Handler.CreateCourse)
	e.PUT("/course/:id", Handler.UpdateCourse)
	e.DELETE("/course/:id", Handler.DeleteCourse)

	// Product route
	e.GET("/product", productHandler.GetProduct)
	e.POST("/product-create", productHandler.CreateProduct)
}
