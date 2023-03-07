package test

import (
	"context"
	"dataway/internal/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestPingGRPC(t *testing.T) {
	client, conn := newGRPCClient(t)
	defer conn.Close()

	if _, err := client.Ping(context.Background(), &emptypb.Empty{}); err != nil {
		t.Fatal(err)
	}
}

func TestSubscribeGRPC(t *testing.T) {
	client, conn := newGRPCClient(t)
	defer conn.Close()

	cID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	tID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	pID := uuid.MustParse("12345678-1234-1234-1234-123456789012")
	if _, err := client.Subscribe(context.Background(), &pb.Subscription{
		ConsumerId: pb.UUIDToPb(cID),
		TomId:      pb.UUIDToPb(tID),
		PropertyId: pb.UUIDToPb(pID),
	}); err != nil {
		if s, ok := status.FromError(err); ok {
			t.Log("error code", s.Code())
			t.Log("error message:", s.Message())
		} else {
			t.Fatal(err)
		}
	}
}
