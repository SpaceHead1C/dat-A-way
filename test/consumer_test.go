package test

import (
	"context"
	. "dataway/internal/domain"
	"testing"
	"time"
)

func TestAddConsumer(t *testing.T) {
	mngr := newTestConsumerManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := mngr.Add(ctx, AddConsumerRequest{
		Name: "me",
		Description: "new",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("=== Consumer ===")
	t.Log("ID:", res.ID.String())
	t.Log("queue:", res.Queue)
	t.Log("name:", res.Name)
	t.Log("description:", res.Description)
}
