package broker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KrylixZA/GoRabbitMqBroker/errors"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
	"github.com/KrylixZA/GoRabbitMqBroker/processing"
	"github.com/streadway/amqp"
)

type messageSubscriber struct {
	config       models.SubscriberConfig
	channel      *amqp.Channel
	queue        amqp.Queue
	errorHandler errors.IHandleError
}

func newMessageSubscriber(config models.SubscriberConfig, channel *amqp.Channel) *messageSubscriber {
	subscriber := messageSubscriber{
		config:       config,
		channel:      channel,
		errorHandler: errors.LogErrorHandler{}, //TODO: Inject this.
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
		subscriber.errorHandler.LogError(err, "Error occurred while declaring exchange")
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
		subscriber.errorHandler.LogError(err, "Error occurred while declaring queue")
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
		subscriber.errorHandler.LogError(err, "Error occured while binding queue to exchange")
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
		subscriber.errorHandler.LogError(err, fmt.Sprintf("Error occurred while attempting to setup consumer on channel againt queue %s", subscriber.config.QueueName))
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			distributedMessage := models.DistributedMessage{}
			err = json.Unmarshal(message.Body, &distributedMessage.Data)
			if err != nil {
				message.Nack(false, subscriber.config.RequeueOnNack)
				subscriber.errorHandler.LogWarning(err, "Error occurred while trying to parse message from RabbitMQ to DistributedMessage struct")
			}
			distributedMessage.CorrelationId = message.CorrelationId
			distributedMessage.MessageId = message.MessageId
			distributedMessage.Timestamp = message.Timestamp

			err = handler.HandleMessage(distributedMessage)
			if err != nil {
				message.Nack(false, subscriber.config.RequeueOnNack)
				subscriber.errorHandler.LogWarning(err, "Error occurred while handler was processing message")
			}
			message.Ack(false) //Acknowledge just this message.
		}

		log.Println("Finished processing all messages on queue... ")
		log.Println("[*] Waiting for messages. To exit press CTRL+C")
	}()

	log.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
