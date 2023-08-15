package test

import (
	"context"
	"dataway/internal/domain"
	"testing"
	"time"
)

func TestAddTom(t *testing.T) {
	mngr := newTestTomManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	id, err := mngr.Add(ctx, domain.RegisterTomRequest{Name: "123"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("=== Tom ===")
	t.Log("ID:", id.String())
}
