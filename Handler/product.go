package Handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"simple-rest-go-echo/Models"
	"simple-rest-go-echo/Usecase"
)

type ProductHandler struct {
	IProductUseCase Models.IProductUseCase
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		IProductUseCase: Usecase.NewProductUseCase(db),
	}
}

func (h ProductHandler) GetAllProduct(c echo.Context) error {
	product, err := h.IProductUseCase.GetAllProduct()
	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, product)
}

func (h ProductHandler) CreateProduct(c echo.Context) error {
	request := &Models.Product{}
	if err := c.Bind(request); err != nil {
		c.Echo().Logger.Error(err)
		return echo.ErrBadRequest
	}

	err := h.IProductUseCase.Create(request)
	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Create Successfully!"})
}
