package Handler

import (
	"github.com/labstack/echo/v4"
	"golang-echo-Domain-Driven-Design/Models"
	"golang-echo-Domain-Driven-Design/Usecase"
	"gorm.io/gorm"
	"net/http"
)

type ProductHandler struct {
	IProductUseCase Models.IProductUseCase
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		IProductUseCase: Usecase.NewProductUseCase(db),
	}
}

func (p ProductHandler) GetAllProduct(c echo.Context) error {
	product, err := p.IProductUseCase.GetAllProduct()
	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, product)
}

func (p ProductHandler) GetProduct(c echo.Context) error {
	request := &Models.SingleProductInput{}
	if err := c.Bind(request); err != nil {
		c.Echo().Logger.Error(err)
		return echo.ErrBadRequest
	}

	product, err := p.IProductUseCase.GetProduct(request)
	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, product)
}

func (p ProductHandler) CreateProduct(c echo.Context) error {
	request := &Models.Product{}
	if err := c.Bind(request); err != nil {
		c.Echo().Logger.Error(err)
		return echo.ErrBadRequest
	}

	err := p.IProductUseCase.Create(request)
	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Create Successfully!"})
}

func (p ProductHandler) UpdateProduct(c echo.Context) error {

	// Retrieve ID from the URL parameter
	productId := c.Param("productId")

	request := &Models.Product{}
	if err := c.Bind(request); err != nil {
		c.Echo().Logger.Error(err)
		return echo.ErrBadRequest
	}

	err := p.IProductUseCase.Update(productId, request)
	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Updated Successfully!"})
}

func (p ProductHandler) DeleteProduct(c echo.Context) error {

	// Retrieve ID from the URL parameter
	productId := c.Param("productId")
	println(productId)

	err := p.IProductUseCase.Delete(productId)
	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted Successfully!"})
}
