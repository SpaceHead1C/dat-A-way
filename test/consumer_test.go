package test

import (
	"context"
	. "dataway/internal/domain"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestAddConsumer(t *testing.T) {
	mngr := newTestConsumerManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := mngr.Add(ctx, AddConsumerRequest{
		Name:        "me",
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

func TestUpdateConsumer(t *testing.T) {
	mngr := newTestConsumerManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	description := "dscr"
	res, err := mngr.Update(ctx, UpdConsumerRequest{
		ID:          uuid.MustParse("12345678-1234-1234-1234-123456789012"),
		Description: &description,
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

func TestGetConsumer(t *testing.T) {
	mngr := newTestConsumerManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := mngr.Get(ctx, uuid.MustParse("12345678-1234-1234-1234-123456789012"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("=== Consumer ===")
	t.Log("ID:", res.ID.String())
	t.Log("queue:", res.Queue)
	t.Log("name:", res.Name)
	t.Log("description:", res.Description)
}
