package rmq

import (
	"dataway/pkg/log"
	"dataway/pkg/message_broker/rmq"
)

const contentTypeAppJSON = "application/json"

type Config struct {
	Publisher *rmq.Publisher
	Logger    *log.Logger
}

type Broker struct {
	publisher *rmq.Publisher
	l         *log.Logger
}

func NewBroker(c Config) (*Broker, error) {
	return &Broker{
		publisher: c.Publisher,
		l:         c.Logger,
	}, nil
}
