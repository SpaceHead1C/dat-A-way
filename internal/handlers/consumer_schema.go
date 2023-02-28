package handlers

import (
	"dataway/internal/domain"
	"fmt"
	"github.com/google/uuid"
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

type UpdConsumerRequestSchema struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (s *UpdConsumerRequestSchema) UpdConsumerRequest() (domain.UpdConsumerRequest, error) {
	out := domain.UpdConsumerRequest{
		Name:        s.Name,
		Description: s.Description,
	}
	id, err := uuid.Parse(s.ID)
	if err != nil {
		return out, fmt.Errorf("parse reference type id error: %s", err)
	}
	out.ID = id
	return out, nil
}

type ConsumerResponseSchema struct {
	ID          string `json:"id"`
	Queue       string `json:"queue"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ConsumerToResponseSchema(c domain.Consumer) ConsumerResponseSchema {
	return ConsumerResponseSchema{
		ID:          c.ID.String(),
		Queue:       c.Queue,
		Name:        c.Name,
		Description: c.Description,
	}
}
