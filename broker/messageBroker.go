//Package broker exposes an abstracted and simplified implementation for interacting with a RabbitMQ node or cluster.
//At a very high level, a user will create/read a configuration which defines whether they intend to be a publisher or a subscriber
//The user will pass this configuration to the message broker which will handle the details of connecting to RabbitMQ and defining the queues and exchanges.
//After this, the user provides the necessary interfaces to the broker and it will handle the details of subscribing and reading messages or publishing messages as per the config.
//Known issues can be found on GitHub (https://github.com/KrylixZA/GoRabbitMqBroker/issues).
//This code is licensed under an MIT license.
//Authors: Simon Headley (KrylixZA).
package broker

import (
	"fmt"

	"github.com/KrylixZA/GoRabbitMqBroker/logs"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
	"github.com/KrylixZA/GoRabbitMqBroker/processing"
	"github.com/streadway/amqp"
)

const rmqConnectionStringTemplate = "amqp://%s:%s@%s:5672/%s"

//IMessageBroker exposes an interface through which users can interact with a RabbitMQ broker.
//Publish exposes functionality to publish an instance of the IDistributedMessageInterface to the configured exchange with the given routing key.
//		The routing key can be a direct routing key, or wildcard if the exchange is configured as a Topic based exchange.
//		The distributed message is an implementation of the IDistributedMessage interface.
//Subscribe exposes functionality to consume messages from a RabbitMQ queue.
//		The handler is a delegate to an implementation of the IMessageHandler interface. This has a HandleMessage function which processes the consumed message.
//		The distributed message is an implementation of the IDistributedMessage interface.
//Close provides a simple endpoint to close the channel and the connection from RabbitMQ.
//		This call should, typically, be deferred immediately after calling a constructor.
type IMessageBroker interface {
	Publish(routingKey string, distributedMessage models.IDistributedMessage) error
	Subscribe(handler processing.IMessageHandler, distributedMessage models.IDistributedMessage) error
	Close()
}

type messageBroker struct {
	config     models.Config
	subscriber *messageSubscriber
	publisher  *messagePublisher
	logger     logs.ILogger
	connection *amqp.Connection
	channel    *amqp.Channel
}

//NewMessageSubscriber initializes a message broker with a given subscriber config.
//		This abstracts away the details of how the connection to RabbitMQ is made and how the queues and exchanges are defined.
//		This will not initialize a publisher. As a result, any attempts to publish a message after using this constructor will not succeed.
//It is imperative that any users of this defer a call to Close() therafter.
//ILogger is some implementation of logs.ILogger.
//		By using an interface, the user of this endpoint can inject any implementation of ILogger.
func NewMessageSubscriber(rmqConfig models.Config, logger logs.ILogger) *messageBroker {
	broker := messageBroker{
		logger: logger,
	}

	err := rmqConfig.Validate()
	if err != nil {
		broker.logger.LogError(err, "Validation of the configuration failed")
	}
	broker.config = rmqConfig
	err = broker.connect()
	if err != nil {
		broker.logger.LogError(err, "Failed to connect to RabbitMQ broker")
	}
	err = broker.createChannel()
	if err != nil {
		broker.logger.LogError(err, "Failed to create channel")
	}

	broker.subscriber = newMessageSubscriber(*rmqConfig.SubscriberConfig, broker.channel, logger)
	return &broker
}

//NewMessagePublisher initializes a message broker with a given publisher config.
//		This abstracts away the details of how the connection to RabbitMQ is made and how the exchanges are defined.
//		This will not initialize a subscriber. As a result, any attempts to subscribe to a queue after using this constructor will not succeed.
//It is imperative that any users of this defer a call to Close() therafter.
//ILogger is some implementation of logs.ILogger.
//		By using an interface, the user of this endpoint can inject any implementation of ILogger.
func NewMessagePublisher(rmqConfig models.Config, logger logs.ILogger) *messageBroker {
	broker := messageBroker{
		logger: logger,
	}

	err := rmqConfig.Validate()
	if err != nil {
		broker.logger.LogError(err, "Validation of the configuration failed")
	}
	broker.config = rmqConfig
	err = broker.connect()
	if err != nil {
		broker.logger.LogError(err, "Failed to connect to RabbitMQ broker")
	}
	err = broker.createChannel()
	if err != nil {
		broker.logger.LogError(err, "Failed to create channel")
	}

	broker.publisher = newMessagePublisher(*rmqConfig.PublisherConfig, broker.channel, logger)
	return &broker
}

//NewMessagePublisherSubscriber initializes a messsage broker with a given config.
//		This abstracts away the details of how the connection to RabbitMQ is made and how the queues and exchanges are defined.
//		This constructor should only ever be used if a user of the service needs to consume messages from a queue and publish to an exchange.
//			It won't always be the case, but this will typically be when a subscriber implements IMessageHandler and then publishes to an exchange from the HandleMessage function.
//It is imperative that any users of this defer a call to Close() therafter.
//ILogger is some implementation of logs.ILogger.
//		By using an interface, the user of this endpoint can inject any implementation of ILogger.
func NewMessagePublisherSubscriber(rmqConfig models.Config, logger logs.ILogger) *messageBroker {
	broker := messageBroker{
		logger: logger,
	}

	err := rmqConfig.Validate()
	if err != nil {
		broker.logger.LogError(err, "Validation of the configuration failed")
	}
	broker.config = rmqConfig
	err = broker.connect()
	if err != nil {
		broker.logger.LogError(err, "Failed to connect to RabbitMQ broker")
	}
	err = broker.createChannel()
	if err != nil {
		broker.logger.LogError(err, "Failed to create channel")
	}

	broker.subscriber = newMessageSubscriber(*rmqConfig.SubscriberConfig, broker.channel, logger)
	broker.publisher = newMessagePublisher(*rmqConfig.PublisherConfig, broker.channel, logger)
	return &broker
}

//Subscribe provides an endpoint for users who wish to consume distributed messages.
//The implementation of IMessageHandler must know how to convert a DistributedMessage into their desired struct in order to process the message correctly.
//The message handler's "HandleMessage" function will be called on demand and asynchronously.
func (broker *messageBroker) Subscribe(handler processing.IMessageHandler) error {
	if broker.subscriber == nil {
		broker.logger.LogError(nil, "RabbitMQ broker was not setup as a subscriber. Cannot subscribe...")
	}
	return broker.subscriber.subscribe(handler)
}

//Publish exposes an endpoint for any users who intend to publish a message.
//Any message that is published to RabbitMQ must satisfy the requirements of the IDistributedMessage interface.
//Any further interfaces that extend the contract of IDistributedMessage can be added at the will of the user.
func (broker *messageBroker) Publish(routingKey string, distributedMessage models.IDistributedMessage) error {
	if broker.publisher == nil {
		broker.logger.LogError(nil, "RabbitMQ broker was not setup as a publisher. Cannot publish...")
	}
	return broker.publisher.publish(routingKey, distributedMessage)
}

//Close closes the connection to the RabbitMQ broker.
//		Close will handle the broker's channel destruction and the connection destruction.
//		Call this function as a deffered execution after creating a connection to RabbitMQ.
func (broker *messageBroker) Close() {
	broker.channel.Close()
	broker.connection.Close()
}

func (broker *messageBroker) connect() error {
	connectionString := fmt.Sprintf(
		rmqConnectionStringTemplate,
		broker.config.Username,
		broker.config.Password,
		broker.config.RabbitMqHost,
		broker.config.VirtualHost)

	connection, err := amqp.Dial(connectionString)
	if err != nil {
		return err
	}

	broker.connection = connection
	return nil
}

func (broker *messageBroker) createChannel() error {
	channel, err := broker.connection.Channel()
	if err != nil {
		return err
	}

	broker.channel = channel
	return nil
}
