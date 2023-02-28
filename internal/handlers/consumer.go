package handlers

import (
	"context"
	"dataway/internal/api"
	"dataway/internal/domain"
	"errors"
	"fmt"
	"net/http"
)

func AddConsumer(ctx context.Context, man *api.ConsumerManager, req AddConsumerRequestSchema) (TextResult, error) {
	out := TextResult{Status: http.StatusCreated}
	cons, err := man.Add(ctx, req.AddConsumerRequest())
	if err != nil {
		out.Status = http.StatusInternalServerError
		return out, err
	}
	out.Payload = cons.ID.String()
	return out, nil
}

func UpdateConsumer(ctx context.Context, man *api.ConsumerManager, req UpdConsumerRequestSchema) (Result, error) {
	out := Result{Status: http.StatusNoContent}
	if req.Name == nil {
		out.Status = http.StatusBadRequest
		return out, fmt.Errorf("name %w", domain.ErrExpected)
	}
	if req.Description == nil {
		out.Status = http.StatusBadRequest
		return out, fmt.Errorf("description %w", domain.ErrExpected)
	}
	r, err := req.UpdConsumerRequest()
	if err != nil {
		out.Status = http.StatusBadRequest
		return out, err
	}
	if _, err := man.Update(ctx, r); err != nil {
		out.Status = http.StatusInternalServerError
		if errors.Is(err, domain.ErrNotFound) {
			out.Status = http.StatusNotFound
		}
		return out, err
	}
	return out, nil
}