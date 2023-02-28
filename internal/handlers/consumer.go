package handlers

import (
	"context"
	"dataway/internal/api"
	"dataway/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
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

func PatchConsumer(ctx context.Context, man *api.ConsumerManager, req UpdConsumerRequestSchema) (Result, error) {
	out := Result{Status: http.StatusOK}
	r, err := req.UpdConsumerRequest()
	if err != nil {
		out.Status = http.StatusBadRequest
		return out, err
	}
	consumer, err := man.Update(ctx, r)
	if err != nil {
		out.Status = http.StatusInternalServerError
		if errors.Is(err, domain.ErrNotFound) {
			out.Status = http.StatusNotFound
		}
		return out, err
	}
	b, err := json.Marshal(ConsumerToResponseSchema(*consumer))
	if err != nil {
		out.Status = http.StatusInternalServerError
		return out, err
	}
	out.Payload = b
	return out, nil
}

func GetConsumer(ctx context.Context, man *api.ConsumerManager, id string) (Result, error) {
	out := Result{Status: http.StatusOK}
	cid, err := uuid.Parse(id)
	if err != nil {
		out.Status = http.StatusBadRequest
		out.Payload = []byte(fmt.Sprintf("parse consumer id error: %s", err))
		return out, err
	}
	consumer, err := man.Get(ctx, cid)
	if err != nil {
		out.Status = http.StatusInternalServerError
		if errors.Is(err, domain.ErrNotFound) {
			out.Status = http.StatusNotFound
		}
		return out, err
	}
	b, err := json.Marshal(ConsumerToResponseSchema(*consumer))
	if err != nil {
		out.Status = http.StatusInternalServerError
		return out, err
	}
	out.Payload = b
	return out, nil
}
