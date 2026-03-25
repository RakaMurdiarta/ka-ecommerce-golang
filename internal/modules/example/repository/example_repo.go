package repository

import (
	"context"

	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/models"
)

type IExampleRepo interface {
	Create(ctx context.Context, m *models.Example) error
}
