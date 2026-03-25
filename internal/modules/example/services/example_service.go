package services

import (
	"context"

	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/delivery"
)

type IExampleService interface {
	CreateExample(ctx context.Context, req *delivery.ExampleRequest) (*delivery.ExampleResponse, error)
}
