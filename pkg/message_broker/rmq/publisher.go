package rmq

import (
	"context"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn *Connection
	ch   *channel
	mu   *sync.RWMutex
}

func (p *Publisher) Close() {
	err := p.ch.close()
	if err != nil {
		p.conn.l.Errorf("close publisher error: %s", err)
	}
	go func() {
		p.ch.closeCh <- closeSignal
	}()
}

type Publishing struct {
	p           Publisher
	exchange    string
	routingKeys []string
	body        []byte
	opts        []PublishOption
}

func (p *Publisher) QueueDeclare(opts ...QueueOption) (string, error) {
	var options queueOptions
	for _, opt := range opts {
		opt(&options)
	}
	return p.conn.ch.queueDeclare(options)
}

func (p *Publisher) Publishing(exchange string, routingKeys []string, body []byte, opts ...PublishOption) *Publishing {
	return &Publishing{
		p:           *p,
		exchange:    exchange,
		routingKeys: routingKeys,
		body:        body,
		opts:        opts,
	}
}

func (p *Publisher) PublishingQueue(queue string, body []byte, opts ...PublishOption) *Publishing {
	return &Publishing{
		p:           *p,
		routingKeys: []string{queue},
		body:        body,
		opts:        opts,
	}
}

func (p *Publishing) Publish(ctx context.Context) error {
	var options publishOptions
	for _, opt := range p.opts {
		opt(&options)
	}
	for _, key := range p.routingKeys {
		if err := p.p.ch.publishWithContext(
			ctx,
			p.exchange,
			key,
			options.mandatory,
			options.immediate,
			options.msg(p.body),
		); err != nil {
			return err
		}
	}
	return nil
}

type Headers amqp.Table

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Set(key string, value any) Headers {
	h[key] = value
	return h
}

func (h Headers) Delete(key string) Headers {
	delete(h, key)
	return h
}

func (h Headers) asNativeType() amqp.Table {
	return amqp.Table(h)
}
