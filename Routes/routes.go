package Routes

import (
	"simple-rest-go-echo/Handler"

	"github.com/labstack/echo/v4"
)

// setup for routing
func SetupRoutes(e *echo.Echo) {
	e.GET("/course", Handler.GetCourse)
	e.POST("/course", Handler.CreateCourse)
	e.PUT("/course/:id", Handler.UpdateCourse)
	e.DELETE("/course/:id", Handler.DeleteCourse)
}
