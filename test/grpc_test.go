package test

import (
	"context"
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
