package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"dataway/internal/api"
	"dataway/internal/domain"
	"dataway/pkg/log"
	"dataway/pkg/message_broker/rmq"

	"github.com/google/uuid"
)

const (
	headerTomID      = "tom_id"
	headerPropertyID = "property_id"
	headerConsumerID = "consumer_id"
)

type ConsumeHandlerConfig struct {
	Logger              *log.Logger
	Timeout             time.Duration
	TomManager          *api.TomManager
	SubscriptionManager *api.SubscriptionManager
	Out                 chan<- domain.Delivery
}

func NewConsumeHandler(c ConsumeHandlerConfig) rmq.Handler {
	return func(d rmq.Delivery) (action rmq.Action) {
		if c.Timeout == 0 {
			c.Timeout = time.Second * 2
		}
		if c.Logger == nil {
			c.Logger = log.GlobalLogger()
		}
		action = rmq.NackDiscard
		dType := domain.DeliveryTypeFromCode(d.Type)
		if dType == domain.UnknownDeliveryType {
			return
		}
		tomID, err := headerAsUUID(d, headerTomID)
		if err != nil {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
		tom, err := c.TomManager.Get(ctx, tomID)
		cancel()
		if err != nil {
			if !errors.Is(err, domain.ErrNotFound) {
				c.Logger.Errorf("process message %s error: %s", d.MessageId, err)
			}
			return
		}
		var propertyID *uuid.UUID
		switch dType {
		case domain.DeliveryTypeValue, domain.DeliveryTypeProperty:
			id, err := headerAsUUID(d, headerPropertyID)
			if err != nil {
				return
			}
			propertyID = &id
		case domain.DeliveryTypeRecord, domain.DeliveryTypeRefType:
		default:
			c.Logger.Infof("unexpected delivery type %s of message %s", d.Type, d.MessageId)
			return
		}
		var consumerIDs []uuid.UUID
		if d.HeaderExists(headerConsumerID) {
			consumerID, err := headerAsUUID(d, headerConsumerID)
			if err != nil {
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
			exists, err := c.SubscriptionManager.SubscriberExists(ctx, domain.SubscriberExistanceRequest{
				ConsumerID: consumerID,
				TomID:      tomID,
			})
			cancel()
			if err != nil {
				return
			}
			if exists {
				consumerIDs = append(consumerIDs, consumerID)
			}
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
			consumerIDs, err = c.SubscriptionManager.SubscriberIDs(ctx, domain.SubscriberIDsRequest{
				TomID:      tomID,
				PropertyID: propertyID,
			})
			cancel()
			if err != nil {
				c.Logger.Errorf("process message %s error: %s", d.MessageId, err)
				return
			}
		}

		c.Out <- domain.Delivery{
			Message: domain.Message{
				ID:        d.MessageId,
				Timestamp: d.Timestamp,
				Type:      dType,
				AppID:     tom.Name,
				Body:      d.Body,
			},
			Recipients: consumerIDs,
		}

		action = rmq.Ack
		return
	}
}

func headerAsString(d rmq.Delivery, key string) (string, error) {
	header, ok := d.HeaderValue(key)
	if !ok {
		return "", fmt.Errorf("header %s is missing", key)
	}
	switch header.(type) {
	case string:
		return header.(string), nil
	default:
		return "", fmt.Errorf("unexpected type %T of header %s", header, key)
	}
}

func headerAsUUID(d rmq.Delivery, key string) (uuid.UUID, error) {
	var out uuid.UUID
	header, err := headerAsString(d, key)
	if err != nil {
		return out, err
	}
	out, err = uuid.Parse(header)
	if err != nil {
		return out, fmt.Errorf("header %s parse error: %w", key, err)
	}
	return out, nil
}
