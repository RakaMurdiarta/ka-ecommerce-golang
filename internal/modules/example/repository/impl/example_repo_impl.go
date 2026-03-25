package impl

import (
	"context"

	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/models"
	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/repository"
)

type ExampleRepoImpl struct {
}

func NewExampleRepoImpl() repository.IExampleRepo {
	return &ExampleRepoImpl{}
}

func (e *ExampleRepoImpl) Create(ctx context.Context, m *models.Example) error {
	return nil
}
