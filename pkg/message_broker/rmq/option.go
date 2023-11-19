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
