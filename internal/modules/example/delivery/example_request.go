package delivery

import "github.com/RakaMurdiarta/ka-ecommerce-golang/internal/models"

type ExampleRequest struct {
	ID uint `json:"id"`
}

func (e *ExampleRequest) ToEntity() *models.Example {
	return &models.Example{
		ID: e.ID,
	}
}
