package models

import (
	"errors"

	"github.com/KrylixZA/GoRabbitMqBroker/enums"
)

//Config describes all the shared configurations needed to connect to RabbitMQ.
//Username is the username the code will use to connect to the RabbitMQ Broker and the Virtual Host.
//		This user must have login access to the RabbitMQ Broker.
//		This user must have full permissions to the Virtual Host.
//Password is the password associated to the user.
//RabbitMqHost is the domain name of the RabbitMQ host. This can be a DNS, IP address or "localhost".
//VirtualHost is the name of the Virtual Host in which the queues and exchanges currently/will exist.
//SubscriberConfig & PublisherConflig are pointers to the configurations for the subscribers and/or publisher.
//		These are optional but there must be at least one configuration provided.
//		Anything that connects to RabbitMQ must be at least a publisher or subscriber, or both.
type Config struct {
	Username         string            `json:"username" doc:"The username that your code will use to connect to RabbitMQ"`
	Password         string            `json:"password" doc:"The password associated to the user that your code will use to connect to RabbitMQ"`
	RabbitMqHost     string            `json:"host" doc:"The host URL (without port) that your code will use to connect to RabbitMQ"`
	VirtualHost      string            `json:"vhost" doc:"The Virtual Host where your queue or exchange will exist"`
	SubscriberConfig *SubscriberConfig `json:"subscriberConfig,omitempty" doc:"The Subscriber configuration."`
	PublisherConfig  *PublisherConfig  `json:"publisherConfig,omitempty" doc:"The Publisher configuration."`
}

//SubscriberConfig describes all the configurations needed to connect to RabbitMQ as a subscriber.
//QueueName is the name of the queue to subscribe to.
//ExchangeName is the name of the exchange the queue will be bound to.
//BindingType is the type of binding used to bind the queue to the exchange.
//RoutingKey is the routing key (topic based - so can include wildcards) that binds the queue to the exchange.
//PrefetchCount is the maximum number of messages to be collected from the queue per subscriber connected to the queue.
//		PrefetchCount can be used to perform, effectively, a Round Robbin load balancing impact.
//		It is likely that this will need to be tweaked over time as you understand your system more and more.
//StrictQueueName defines whether the code must validate the queue name, or allow RabbitMQ to generate it's own queue name that has no meaning to us.
//Durable defines whether or not RabbitMQ should persist messages to cache/disk if they are not acknowledged in the event of a crash or restart of the RabbitMQ server.
//AutoDeleteQueue defines whether the queue should be automatically deleted or not when there are no more subscribers to the queue.
//RequeueOnNack defines whether or not the message should be requeued in the event of an error while trying to process the message. The default is false.
//		Override this if you want messages to be replayed until they pass (can potentially bottleneck the queueing by causing errors).
type SubscriberConfig struct {
	QueueName       string            `json:"queueName" doc:"The name of the queue to subscribe to"`
	ExchangeName    string            `json:"exchangeName" doc:"The name of the exchange the queue is bound to"`
	BindingType     enums.BindingType `json:"bindingType,int" doc:"The type of binding the queue should use when binding to the queue. Default is fanout"`
	RoutingKey      string            `json:"routingKey" doc:"The routing key that binds the queeu to the exchange"`
	PrefetchCount   int               `json:"prefetchCount" doc:"The maximum amount of messages to consume at once"`
	StrictQueueName bool              `json:"strictQueueName" doc:"Set to true if queue names must be defined. If false, RabbitMQ will auto-generate queue names. Default value is false"`
	Durable         bool              `json:"durable" doc:"Set to true if RabbitMQ should persist the messages to cache/disk if they are not acknowledged in the event of a crash or restart. Default is false"`
	AutoDeleteQueue bool              `json:"autoDeleteQueue" doc:"Set to true if the queue should be deleted automatically as soon as there are no more subscribers. Default value is false"`
	RequeueOnNack   bool              `json:"requeueOnNack" doc:"Set to true if messages should be requeued when they are nacked. Default is false"`
}

//PublisherConfig describes all the configurations needed to connect to RabbitMQ as a publisher.
//ExchangeName is the name that the publisher will publisher to.
//		The routing key used is determined during runtime when calling the message broker's publish function.
//BindingType is the type of binding used to bind any queue to the exchange.
//Durable defines whether or not RabbitMQ should persist messages to cache/disk if they are not acknowledged in the event of a crash or restart of the RabbitMQ server.
//MandatoryQueueBind is a condition set when publishing to know if a queue is bound to the exchange. If this is set to true, and no queue is bound, publishing will fail.
type PublisherConfig struct {
	ExchangeName       string            `json:"exchangeName" doc:"The exchange to publish to"`
	BindingType        enums.BindingType `json:"bindingType,int" doc:"The type of binding the queue should use when binding to the queue. Default is fanout"`
	Durable            bool              `json:"durable" doc:"Set to true if RabbitMQ should persist the messages to cache/disk if they are not acknowledged in the event of a crash or restart. Default is false"`
	MandatoryQueueBind bool              `json:"mandatoryQueueBind" doc:"Set to true if a queue must be bound to the queue for publishing to be successful. Default is false."`
}

//Validate enforces that the configuration provided to the messageBroker is all well-formed & correct.
//		Validate will enforce all the necessary connection string properties are present.
//		Validate will also enforce that a consumer of the RabbitMQ broker is at least a producer or a consumer (it could be both)
func (config *Config) Validate() error {
	if config.Username == "" {
		return errors.New("username is empty string")
	}
	if config.Password == "" {
		return errors.New("password is empty string")
	}
	if config.RabbitMqHost == "" {
		return errors.New("host is empty string")
	}
	if config.VirtualHost == "" {
		return errors.New("vhost is empty string")
	}
	if config.VirtualHost == "/" {
		config.VirtualHost = ""
	}
	if config.SubscriberConfig == nil && config.PublisherConfig == nil {
		return errors.New("subscriberConfig and publisherConfig are missing. A consumer of the RabbitMQ broker must be a producer, or a consumer, or both")
	}
	if config.SubscriberConfig != nil {
		return config.SubscriberConfig.Validate()
	}
	if config.PublisherConfig != nil {
		return config.PublisherConfig.Validate()
	}

	return nil
}

//Validate enforces that the subscriber configuration provided is all well-formed & correct.
//		Validate will enforce that if strictQueueName is true, a queue name is provided.
//		Validate will enforce that an exchange name is provided to which the queue will be bound.
//		Validate will enforce that if the Binding Type is Direct or Topic, a routing key is provided.
func (config *SubscriberConfig) Validate() error {
	if config.StrictQueueName && config.QueueName == "" {
		return errors.New("subscriberConfig.strictQueueName is set to true but subscriberConfig.queueName is empty string. If you wish to use auto-generated queue names, set strictQueueName to false")
	}
	if config.ExchangeName == "" {
		return errors.New("subscriberConfig.exchangeName is empty string. Although RabbitMQ allows for auto-generating exchange names, it becomes complex to manage when binding queues. As such, we force an exchangeName to be supplied in the config")
	}
	if config.BindingType < 0 || config.BindingType > 2 {
		return errors.New("subscriberConfig.bindingType is out of range. Acceptable options are 0 = Fanout, 1 = Direct, 2 = Topic")
	}
	if (config.BindingType == enums.Direct || config.BindingType == enums.Topic) && config.RoutingKey == "" {
		return errors.New("subscriberConfig.routingKey is empty string. Cannot use an empty routing key to bind a queue to an exchange when using Direct or Topic based routing")
	}
	if config.PrefetchCount < 0 {
		return errors.New("subscriberConfig.prefetchCount cannot be less than zero")
	}

	return nil
}

//Validate enforces that the publisher configuration provided is all well-formed & correct.
//		Validate will enforce that an exchange name is provided.
func (config *PublisherConfig) Validate() error {
	if config.ExchangeName == "" {
		return errors.New("publisherConfig.exchangeName is empty string. Although RabbitMQ allows for auto-generating exchange names, it becomes complex to manage when binding queues. As such, we force an exchangeName to be supplied in the config")
	}
	if config.BindingType < 0 || config.BindingType > 2 {
		return errors.New("publisherConfig.bindingType is out of range. Acceptable options are 0 = Fanout, 1 = Direct, 2 = Topic")
	}

	return nil
}
