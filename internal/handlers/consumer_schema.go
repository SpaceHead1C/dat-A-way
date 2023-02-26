package handlers

import (
	"dataway/internal/domain"
)

type AddConsumerRequestSchema struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s AddConsumerRequestSchema) AddConsumerRequest() domain.AddConsumerRequest {
	return domain.AddConsumerRequest{
		Name:        s.Name,
		Description: s.Description,
	}
}
