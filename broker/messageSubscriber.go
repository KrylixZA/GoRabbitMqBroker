package broker

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/KrylixZA/GoRabbitMqBroker/logs"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
	"github.com/KrylixZA/GoRabbitMqBroker/processing"
	"github.com/streadway/amqp"
)

type messageSubscriber struct {
	config  models.SubscriberConfig
	channel *amqp.Channel
	queue   amqp.Queue
	logger  logs.ILogger
}

func newMessageSubscriber(config models.SubscriberConfig, channel *amqp.Channel, logger logs.ILogger) *messageSubscriber {
	subscriber := messageSubscriber{
		config:  config,
		channel: channel,
		logger:  logger,
	}

	//Declare the exchange
	err := channel.ExchangeDeclare(
		config.ExchangeName,
		config.BindingType.String(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		subscriber.logger.LogError(err, "Error occurred while declaring exchange")
	}

	//Declare the queue
	q, err := channel.QueueDeclare(
		config.QueueName,
		config.Durable,
		config.AutoDeleteQueue,
		false,
		false,
		nil,
	)
	if err != nil {
		subscriber.logger.LogError(err, "Error occurred while declaring queue")
	}
	subscriber.queue = q

	//Set the prefetch count
	channel.Qos(
		config.PrefetchCount,
		0,
		false,
	)

	//Bind queue to exchange
	err = channel.QueueBind(
		q.Name,
		config.RoutingKey,
		config.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		subscriber.logger.LogError(err, "Error occured while binding queue to exchange")
	}

	return &subscriber
}

func (subscriber *messageSubscriber) subscribe(handler processing.IMessageHandler) error {
	messages, err := subscriber.channel.Consume(
		subscriber.queue.Name,
		"", //TODO: Change this consumer name to something real - service name etc.
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		subscriber.logger.LogError(err, fmt.Sprintf("Error occurred while attempting to setup consumer on channel againt queue %s", subscriber.config.QueueName))
	}

	wg := sync.WaitGroup{}

	for message := range messages {
		wg.Add(1)
		go func(message amqp.Delivery) {
			defer wg.Done()
			distributedMessage := models.DistributedMessage{}
			err := json.Unmarshal(message.Body, &distributedMessage.Data)
			if err != nil {
				message.Nack(false, subscriber.config.RequeueOnNack)
				subscriber.logger.LogWarning(fmt.Sprintf("Error occurred while trying to parse message from RabbitMQ to DistributedMessage struct\n\n%s",
					err))
			}
			distributedMessage.CorrelationId = message.CorrelationId
			distributedMessage.MessageId = message.MessageId
			distributedMessage.Timestamp = message.Timestamp

			err = handler.HandleMessage(distributedMessage)
			if err != nil {
				message.Nack(false, subscriber.config.RequeueOnNack)
				subscriber.logger.LogWarning(fmt.Sprintf("Error occurred while handler was processing message\n\n%s",
					err))
			}
			message.Ack(false) //Acknowledge just this message.
		}(message)
	}
	wg.Wait()

	return nil
}
