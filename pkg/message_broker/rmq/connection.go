package rmq

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"dataway/pkg/log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	defaultPort uint16 = 5672

	defaultUser     = "guest"
	defaultPassword = "guest"
)

type ConnectionConfig struct {
	Logger   *log.Logger
	Host     string
	Port     uint16
	User     string
	Password string
	VHost    string
}

type Connection struct {
	url  string
	conn *amqp.Connection
	ch   *channel
	mu   *sync.RWMutex
	l    *log.Logger
}

func New(c ConnectionConfig) (*Connection, error) {
	if c.Host == "" {
		return nil, errors.New("address can not be empty")
	}
	if c.Port == 0 {
		c.Port = defaultPort
	}
	if c.Logger == nil {
		c.Logger = log.GlobalLogger()
	}
	if strings.TrimSpace(c.User) == "" {
		c.User = defaultUser
	}
	if strings.TrimSpace(c.Password) == "" {
		c.Password = defaultPassword
	}
	url := connectionString(c)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	out := &Connection{
		url:  url,
		conn: conn,
		mu:   &sync.RWMutex{},
		l:    c.Logger,
	}
	out.ch, err = out.channel()
	if err != nil {
		return nil, err
	}
	go out.startNotifyClose()
	return out, nil
}

func (conn *Connection) Close() {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if err := conn.conn.Close(); err != nil {
		conn.l.Errorf("connection close error: %s", err)
	}
}

func (conn *Connection) Consumer(queue string, handler Handler, opts ...ConsumeOption) (*Consumer, error) {
	ch, err := conn.channel()
	if err != nil {
		return nil, err
	}
	var options consumeOptions
	for _, opt := range opts {
		opt(&options)
	}
	out := &Consumer{
		conn: conn,
		ch:   ch,
		mu:   &sync.RWMutex{},

		isClosedMu: &sync.RWMutex{},

		queue:   queue,
		handler: handler,
		opts:    options,
	}
	go func() {
		for err := range out.ch.errCh {
			conn.l.Infof("successful consumer recovery from: %s", err)
			if err := out.start(); err != nil {
				conn.l.Errorf("start consuming error: %s", err)
				return
			}
		}
	}()
	if err := out.start(); err != nil {
		return nil, fmt.Errorf("start consuming error: %w", err)
	}
	return out, nil
}

func connectionString(c ConnectionConfig) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.VHost)
}

func (conn *Connection) channel() (*channel, error) {
	ch, err := conn.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	closeCh := make(chan struct{})
	errCh := make(chan error)
	go func() {
		<-closeCh
		conn.l.Info("closed")
		close(errCh)
	}()
	out := &channel{
		conn:    conn,
		ch:      ch,
		errCh:   errCh,
		closeCh: closeCh,
		mu:      &sync.RWMutex{},
	}
	go out.listenNotifies()
	return out, nil
}

func (conn *Connection) startNotifyClose() {
	notifyCloseChan := conn.conn.NotifyClose(make(chan *amqp.Error, 1))
	err := <-notifyCloseChan
	if err != nil {
		conn.l.Errorf("connection close with error: %s\nReconnecting...", err)
		conn.reconnecting()
		conn.l.Info("successfully reconnected to amqp server")
	}
}

func (conn *Connection) reconnecting() {
	for {
		time.Sleep(time.Second)
		err := conn.reconnect()
		if err == nil {
			go conn.startNotifyClose()
			return
		}
		conn.l.Error(err.Error())
	}
}

func (conn *Connection) reconnect() error {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	newConn, err := amqp.Dial(conn.url)
	if err != nil {
		return err
	}
	if err = conn.conn.Close(); err != nil {
		conn.l.Errorf("close channel error: %s", err)
	}
	conn.conn = newConn
	return nil
}
