package rmq

import (
	"context"
	"errors"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var closeSignal = struct{}{}

type channel struct {
	conn    *Connection
	ch      *amqp.Channel
	errCh   chan error
	closeCh chan struct{}
	mu      *sync.RWMutex
}

func (ch *channel) close() error {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	return ch.ch.Close()
}

func (ch *channel) listenNotifies() {
	closeCh := ch.ch.NotifyClose(make(chan *amqp.Error, 1))
	cancelCh := ch.ch.NotifyCancel(make(chan string, 1))
	var err error
	select {
	case err = <-closeCh:
		if err == nil {
			return
		}
	case msg := <-cancelCh:
		err = errors.New(msg)
	}
	ch.reconnecting()
	ch.errCh <- err
}

func (ch *channel) reconnecting() {
	for {
		time.Sleep(time.Second)
		err := ch.reconnect()
		if err == nil {
			go ch.listenNotifies()
			return
		}
		ch.conn.l.Error(err.Error())
	}
}

func (ch *channel) reconnect() error {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	newCh, err := ch.conn.conn.Channel()
	if err != nil {
		return err
	}
	if err = ch.ch.Close(); err != nil {
		ch.conn.l.Errorf("close channel error: %s", err)
	}
	ch.ch = newCh
	return nil
}

func (ch *channel) queueDeclare(opts queueOptions) (string, error) {
	ch.mu.RLock()
	defer ch.mu.RUnlock()

	f := ch.ch.QueueDeclare
	if opts.passive {
		f = ch.ch.QueueDeclarePassive
	}
	q, err := f(
		opts.name,
		opts.durable,
		opts.autoDelete,
		opts.exclusive,
		opts.noWait,
		opts.args.asNativeType(),
	)
	if err != nil {
		return "", err
	}
	return q.Name, nil
}

func (ch *channel) publishWithContext(ctx context.Context, exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error {
	ch.mu.RLock()
	defer ch.mu.RUnlock()
	return ch.ch.PublishWithContext(ctx, exchange, key, mandatory, immediate, msg)
}

func (ch *channel) consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	ch.mu.RLock()
	defer ch.mu.RUnlock()
	return ch.ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}
