package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type DeliveryType uint

const (
	UnknownDeliveryType DeliveryType = iota
	DeliveryTypeRecord
	DeliveryTypeProperty
	DeliveryTypeRefType
	DeliveryTypeValue
)

func (dt DeliveryType) Code() string {
	switch dt {
	case DeliveryTypeValue:
		return "value"
	case DeliveryTypeRecord:
		return "record"
	case DeliveryTypeProperty:
		return "property"
	case DeliveryTypeRefType:
		return "reference_type"
	default:
		return ""
	}
}

func (dt DeliveryType) String() string {
	out := dt.Code()
	if out == "" {
		out = "unknown"
	}
	return out
}

func DeliveryTypeFromCode(code string) DeliveryType {
	switch code {
	case "value":
		return DeliveryTypeValue
	case "record":
		return DeliveryTypeRecord
	case "property":
		return DeliveryTypeProperty
	case "reference_type":
		return DeliveryTypeRefType
	default:
		return UnknownDeliveryType
	}
}

type DeliveryBroker interface {
	SendDelivery(context.Context, uuid.UUID, Message) error
}

type Delivery struct {
	Message
	Recipients []uuid.UUID
}

type Message struct {
	ID        string
	Timestamp time.Time
	Type      DeliveryType
	AppID     string
	Body      []byte
}
