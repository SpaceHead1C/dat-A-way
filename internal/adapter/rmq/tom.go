package rmq

import (
	"context"

	"dataway/internal/domain"
	pkgrmq "dataway/pkg/message_broker/rmq"
)

func (b *Broker) DeclareTomQueue(ctx context.Context, tom domain.Tom) error {
	_, err := b.publisher.QueueDeclare(
		pkgrmq.QueueWithName(tom.QueueName()),
		pkgrmq.QueueIsDurable(),
		pkgrmq.QueueWithArgs(pkgrmq.NewQueueArgs().AsClassic()),
	)
	return err
}
