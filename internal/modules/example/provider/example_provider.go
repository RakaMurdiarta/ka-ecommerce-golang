package provider

import (
	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/handler"
	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func NewExampleProvider(privateRoute *echo.Group, publicRoute *echo.Group, es services.IExampleService) {

	v := validator.New()

	exampleHandler := handler.NewExampleHandler(es, v)

	exampleRoute := publicRoute.Group("/example")

	exampleRoute.POST("", exampleHandler.Create)

}
