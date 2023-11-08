package Handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-rest-go-echo/Models"
	"simple-rest-go-echo/Usecase"
)

type ProductHandler struct {
	IProductUseCase Models.IProductUseCase
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		IProductUseCase: Usecase.NewProductUseCase(),
	}
}

func (h ProductHandler) GetProduct(c echo.Context) error {
	return nil
}

func (h ProductHandler) CreateProduct(c echo.Context) error {
	request := &[]Models.Product{}
	if err := c.Bind(request); err != nil {
		c.Echo().Logger.Error(err)
		return echo.ErrBadRequest
	}

	err := h.IProductUseCase.Create(request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusUnprocessableEntity, err)
}
