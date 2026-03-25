package delivery

import "github.com/RakaMurdiarta/ka-ecommerce-golang/internal/models"

type ExampleResponse struct {
	ID uint `json:"id"`
}

func ToResponse(m *models.Example) *ExampleResponse {
	return &ExampleResponse{
		ID: m.ID,
	}
}
