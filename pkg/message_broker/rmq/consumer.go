package rmq

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *Connection
	ch   *channel
	mu   *sync.RWMutex

	isClosedMu *sync.RWMutex
	isClosed   bool

	queue   string
	handler Handler
	opts    consumeOptions
}

func (cons *Consumer) Close() {
	cons.isClosedMu.Lock()
	defer cons.isClosedMu.Unlock()
	cons.isClosed = true
	err := cons.ch.close()
	if err != nil {
		cons.conn.l.Errorf("close consumer's channel error: %s", err)
	}
	go func() {
		cons.ch.closeCh <- closeSignal
	}()
}

func (cons *Consumer) start() error {
	msgs, err := cons.ch.consume(
		cons.queue,
		cons.opts.name,
		cons.opts.autoAck,
		cons.opts.exclusive,
		cons.opts.noLocal,
		cons.opts.noWait,
		cons.opts.args.asNativeType(),
	)
	if err != nil {
		return err
	}
	go cons.handle(msgs, cons.opts.autoAck, cons.handler)
	return nil
}

func (cons *Consumer) handle(msgs <-chan amqp.Delivery, autoAck bool, handler Handler) {
	for msg := range msgs {
		if cons.getIsClosed() {
			break
		}

		res := handler(Delivery{msg})
		if autoAck {
			continue
		}
		switch res {
		case Ack:
			err := msg.Ack(false)
			if err != nil {
				cons.conn.l.Errorf("ack message error: %s", err)
			}
		case NackDiscard:
			err := msg.Nack(false, false)
			if err != nil {
				cons.conn.l.Errorf("nack message error: %s", err)
			}
		case NackRequeue:
			err := msg.Nack(false, true)
			if err != nil {
				cons.conn.l.Errorf("nack message error: %s", err)
			}
		}
	}
}

func (cons *Consumer) getIsClosed() bool {
	cons.isClosedMu.RLock()
	defer cons.isClosedMu.RUnlock()
	return cons.isClosed
}

type Action uint

const (
	Ack Action = iota
	NackDiscard
	NackRequeue
)

type Delivery struct {
	amqp.Delivery
}

func (d Delivery) HeaderExists(key string) bool {
	_, ok := d.Headers[key]
	return ok
}

func (d Delivery) HeaderValue(key string) (any, bool) {
	value, ok := d.Headers[key]
	return value, ok
}

type Handler func(d Delivery) (action Action)

type ConsumeArgs amqp.Table

func NewConsumeArgs() ConsumeArgs {
	return make(ConsumeArgs)
}

func (ca ConsumeArgs) SetArg(key string, value any) ConsumeArgs {
	ca[key] = value
	return ca
}

func (ca ConsumeArgs) DeleteArg(key string) ConsumeArgs {
	delete(ca, key)
	return ca
}

func (ca ConsumeArgs) asNativeType() amqp.Table {
	return amqp.Table(ca)
}
