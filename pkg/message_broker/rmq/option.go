package rmq

type QueueOption func(options *queueOptions)

type queueOptions struct {
	name       string
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
	passive    bool
	args       QueueArgs
}

func QueueWithName(name string) QueueOption {
	return func(options *queueOptions) {
		options.name = name
	}
}

func QueueIsDurable() QueueOption {
	return func(options *queueOptions) {
		options.durable = true
	}
}

func QueueIsAutoDelete() QueueOption {
	return func(options *queueOptions) {
		options.autoDelete = true
	}
}

func QueueIsExclusive() QueueOption {
	return func(options *queueOptions) {
		options.exclusive = true
	}
}

func QueueIsNoWait() QueueOption {
	return func(options *queueOptions) {
		options.noWait = true
	}
}

func QueuePassiveDeclare() QueueOption {
	return func(options *queueOptions) {
		options.passive = true
	}
}

func QueueWithArgs(args QueueArgs) QueueOption {
	return func(options *queueOptions) {
		options.args = args
	}
}

type publishOptions struct {
	headers         Headers
	contentType     string
	contentEncoding string
	deliveryMode    uint8
	priority        uint8
	correlationID   string
	replyTo         string
	expiration      string
	messageID       string
	timestamp       time.Time
	messageType     string
	userID          string
	appID           string
	mandatory       bool
	immediate       bool
}

func (opts publishOptions) msg(body []byte) amqp.Publishing {
	return amqp.Publishing{
		Headers:         opts.headers.asNativeType(),
		ContentType:     opts.contentType,
		ContentEncoding: opts.contentEncoding,
		DeliveryMode:    opts.deliveryMode,
		Priority:        opts.priority,
		CorrelationId:   opts.correlationID,
		ReplyTo:         opts.replyTo,
		Expiration:      opts.expiration,
		MessageId:       opts.messageID,
		Timestamp:       opts.timestamp,
		Type:            opts.messageType,
		UserId:          opts.userID,
		AppId:           opts.appID,
		Body:            body,
	}
}

type PublishOption func(*publishOptions)

// WithPublishOptionHeaders sets application or exchange specific fields
func WithPublishOptionHeaders(headers Headers) PublishOption {
	return func(opts *publishOptions) {
		opts.headers = headers
	}
}

// WithPublishOptionContentType sets MIME content type
func WithPublishOptionContentType(contentType string) PublishOption {
	return func(opts *publishOptions) {
		opts.contentType = contentType
	}
}

// WithPublishOptionContentEncoding sets MIME content encoding
func WithPublishOptionContentEncoding(contentEncoding string) PublishOption {
	return func(opts *publishOptions) {
		opts.contentEncoding = contentEncoding
	}
}

// WithPublishOptionPersistent sets delivery mode as Persistent = 2 (default mode is Transient = 0 or 1)
func WithPublishOptionPersistent() PublishOption {
	return func(opts *publishOptions) {
		opts.deliveryMode = amqp.Persistent
	}
}

// WithPublishOptionPriority sets priority in range 0 to 9
func WithPublishOptionPriority(priority uint8) PublishOption {
	return func(opts *publishOptions) {
		if priority > 9 {
			priority = 9
		}
		opts.priority = priority
	}
}

// WithPublishOptionCorrelationID sets correlation identifier
func WithPublishOptionCorrelationID(correlationID string) PublishOption {
	return func(opts *publishOptions) {
		opts.correlationID = correlationID
	}
}

// WithPublishOptionReplyTo sets address to to reply to (ex: RPC)
func WithPublishOptionReplyTo(replyTo string) PublishOption {
	return func(opts *publishOptions) {
		opts.replyTo = replyTo
	}
}

// WithPublishOptionExpiration sets message expiration spec (string value in milliseconds)
func WithPublishOptionExpiration(expiration string) PublishOption {
	return func(opts *publishOptions) {
		opts.expiration = expiration
	}
}

// WithPublishOptionMessageID sets message identifier
func WithPublishOptionMessageID(messageID string) PublishOption {
	return func(opts *publishOptions) {
		opts.messageID = messageID
	}
}

// WithPublishOptionTimestamp sets message timestamp
func WithPublishOptionTimestamp(timestamp time.Time) PublishOption {
	return func(opts *publishOptions) {
		opts.timestamp = timestamp
	}
}

// WithPublishOptionType sets message type name
func WithPublishOptionType(messageType string) PublishOption {
	return func(opts *publishOptions) {
		opts.messageType = messageType
	}
}

// WithPublishOptionUserID sets user identifier - ex: "guest"
func WithPublishOptionUserID(userID string) PublishOption {
	return func(opts *publishOptions) {
		opts.userID = userID
	}
}

// WithPublishOptionAppID sets application identifier
func WithPublishOptionAppID(appID string) PublishOption {
	return func(opts *publishOptions) {
		opts.appID = appID
	}
}

/*
WithPublishOptionMandatory sets mode when message will be sent back on the channel for undeliverable messages
if no queue is bound
*/
func WithPublishOptionMandatory() PublishOption {
	return func(opts *publishOptions) {
		opts.mandatory = true
	}
}

/*
WithPublishOptionImmediate sets mode when message will be sent back on the channel for undeliverable messages
if no consumer on the matched queue is ready to accept the delivery
*/
func WithPublishOptionImmediate() PublishOption {
	return func(opts *publishOptions) {
		opts.immediate = true
	}
}

type consumeOptions struct {
	queueOpts queueOptions
	name      string
	autoAck   bool
	exclusive bool
	noWait    bool
	noLocal   bool
	args      ConsumeArgs
}

type ConsumeOption func(opts *consumeOptions)

// WithConsumeOptionConsumerName sets consumer's name
func WithConsumeOptionConsumerName(name string) ConsumeOption {
	return func(opts *consumeOptions) {
		opts.name = name
	}
}

// WithConsumeOptionAutoAck sets auto acknowledge mode on
func WithConsumeOptionAutoAck() ConsumeOption {
	return func(opts *consumeOptions) {
		opts.autoAck = true
	}
}

// WithConsumeOptionExclusive sets exclusive mode on for the consumer
func WithConsumeOptionExclusive() ConsumeOption {
	return func(opts *consumeOptions) {
		opts.exclusive = true
	}
}

// WithConsumeOptionNoWait sets no wait mode on
func WithConsumeOptionNoWait() ConsumeOption {
	return func(opts *consumeOptions) {
		opts.noWait = true
	}
}

// WithConsumeOptionNoLocal sets no local mode on
func WithConsumeOptionNoLocal() ConsumeOption {
	return func(opts *consumeOptions) {
		opts.noLocal = true
	}
}

// WithConsumeOptionArguments sets consuming arguments
func WithConsumeOptionArguments(args ConsumeArgs) ConsumeOption {
	return func(opts *consumeOptions) {
		opts.args = args
	}
}
