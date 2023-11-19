package rmq

import amqp "github.com/rabbitmq/amqp091-go"

const queueDLEArg = "x-dead-letter-exchange"

type QueueArgs amqp.Table

func NewQueueArgs() QueueArgs {
	return make(QueueArgs)
}

func (qa QueueArgs) SetArg(key string, value any) QueueArgs {
	qa[key] = value
	return qa
}

func (qa QueueArgs) DeleteArg(key string) QueueArgs {
	delete(qa, key)
	return qa
}

func (qa QueueArgs) SetDLEArg(value string) QueueArgs {
	return qa.SetArg(queueDLEArg, value)
}

func (qa QueueArgs) SetTypeArg(value string) QueueArgs {
	return qa.SetArg(amqp.QueueTypeArg, value)
}

func (qa QueueArgs) AsClassic() QueueArgs {
	return qa.SetTypeArg(amqp.QueueTypeClassic)
}

func (qa QueueArgs) asNativeType() amqp.Table {
	return amqp.Table(qa)
}
