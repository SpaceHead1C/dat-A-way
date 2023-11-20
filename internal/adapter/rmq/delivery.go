package rmq

import (
	"context"

	"dataway/internal/domain"
	"dataway/pkg/message_broker/rmq"

	"github.com/google/uuid"
)

func (b *Broker) SendDelivery(ctx context.Context, consumerID uuid.UUID, m domain.Message) error {
	return b.publisher.PublishingQueue(
		domain.ConsumerQueue(consumerID),
		m.Body,
		rmq.WithPublishOptionPersistent(),
		rmq.WithPublishOptionContentType(contentTypeAppJSON),
		rmq.WithPublishOptionMessageID(m.ID),
		rmq.WithPublishOptionType(m.Type.Code()),
		rmq.WithPublishOptionAppID(m.AppID),
		rmq.WithPublishOptionTimestamp(m.Timestamp),
	).Publish(ctx)
}
