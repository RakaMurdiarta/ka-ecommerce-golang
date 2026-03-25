package impl

import (
	"context"
	"fmt"

	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/delivery"
	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/repository"
	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/services"
)

type ExampleServiceImpl struct {
	er repository.IExampleRepo
}

func NewExampleService(er repository.IExampleRepo) services.IExampleService {
	return &ExampleServiceImpl{
		er: er,
	}
}

func (es *ExampleServiceImpl) CreateExample(ctx context.Context, req *delivery.ExampleRequest) (*delivery.ExampleResponse, error) {

	exampple := req.ToEntity()
	err := es.er.Create(ctx, exampple)

	if err != nil {
		return nil, fmt.Errorf("error create example")
	}

	response := delivery.ToResponse(exampple)

	return response, nil

}
