package test

import (
	"context"
	"testing"
	"time"
)

func TestAddTom(t *testing.T) {
	mngr := newTestTomManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	id, err := mngr.Add(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("=== Tom ===")
	t.Log("ID:", id.String())
}
