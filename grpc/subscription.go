package grpc

import (
	"context"
	"dataway/internal/api"
	"dataway/internal/domain"
	"dataway/internal/pb"
	"dataway/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Subscribe(ctx context.Context, man *api.SubscriptionManager, req *pb.Subscription) (*pb.Subscription, error) {
	s, err := subscriptionFromPb(req)
	if err != nil {
		return nil, err
	}
	out, err := man.Add(ctx, domain.AddSubscriptionRequest{
		ConsumerID: s.ConsumerID,
		TomID:      s.TomID,
		PropertyID: s.PropertyID,
	})
	if err != nil {
		log.LoggerFromContext(ctx).Errorf("add subscription error: %s", err)
		return nil, status.Errorf(codes.Internal, "add subscription error")
	}
	return subscriptionToPb(*out), nil
}

func DeleteSubscription(ctx context.Context, man *api.SubscriptionManager, req *pb.Subscription) error {
	s, err := subscriptionFromPb(req)
	if err != nil {
		return err
	}
	if _, err := man.Delete(ctx, domain.DeleteSubscriptionRequest{
		ConsumerID: s.ConsumerID,
		TomID:      s.TomID,
		PropertyID: s.PropertyID,
	}); err != nil {
		log.LoggerFromContext(ctx).Errorf("delete subscription error: %s", err)
		return status.Errorf(codes.Internal, "delete subscription error")
	}
	return nil
}

func subscriptionFromPb(in *pb.Subscription) (domain.Subscription, error) {
	var out domain.Subscription
	cID, err := pb.UUIDFromPb(in.ConsumerId)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, "parse consumer ID error: %s", err)
	}
	out.ConsumerID = cID
	tID, err := pb.UUIDFromPb(in.TomId)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, "parse tom ID error: %s", err)
	}
	out.TomID = tID
	pID, err := pb.UUIDFromPb(in.PropertyId)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, "parse property ID error: %s", err)
	}
	out.PropertyID = pID
	return out, nil
}

func subscriptionToPb(in domain.Subscription) *pb.Subscription {
	return &pb.Subscription{
		ConsumerId: pb.UUIDToPb(in.ConsumerID),
		TomId:      pb.UUIDToPb(in.TomID),
		PropertyId: pb.UUIDToPb(in.PropertyID),
	}
}
