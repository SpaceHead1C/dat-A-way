package handlers

import (
	"context"
	"dataway/internal/api"
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
