package handler

import (
	"net/http"

	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/delivery"
	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type exampleHandler struct {
	es services.IExampleService
	v  *validator.Validate
}

func NewExampleHandler(
	es services.IExampleService,
	v *validator.Validate,
) *exampleHandler {
	return &exampleHandler{
		es: es,
		v:  v,
	}
}

func (e *exampleHandler) Create(c *echo.Context) error {
	req := new(delivery.ExampleRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed",
		})
	}

	if err := e.v.StructCtx(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed",
		})
	}

	example, err := e.es.CreateExample(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed",
		})
	}

	return c.JSON(http.StatusOK, map[string]*delivery.ExampleResponse{
		"data": example,
	})
}
