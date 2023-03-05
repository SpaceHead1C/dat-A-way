package test

import (
	"context"
	"dataway/internal/domain"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestAddSubscription(t *testing.T) {
	mngr := newTestSubscriptionManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := mngr.Add(ctx, domain.AddSubscriptionRequest{
		ConsumerID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		TomID:      uuid.MustParse("22222222-2222-2222-2222-222222222222"),
		PropertyID: uuid.MustParse("12345678-1234-1234-1234-123456789012"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("=== Subscription ===")
	t.Log("consumer ID:", res.ConsumerID.String())
	t.Log("tom ID:", res.TomID.String())
	t.Log("property ID:", res.PropertyID.String())
}
