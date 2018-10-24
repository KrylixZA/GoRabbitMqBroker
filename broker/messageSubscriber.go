package broker

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

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
		true,
		false,
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
		config.QueueName,
		config.RoutingKey,
		config.ExchangeName,
		false,
		nil,
	)

	return &subscriber
}

func (subscriber *messageSubscriber) subscribe(handler processing.IMessageHandler, distributedMessage models.IDistributedMessage) error {
	messages, err := subscriber.channel.Consume(
		subscriber.config.QueueName,
		"", //TODO: Change this consumer name to something real - service name etc.
		false,
		false,
		false,
		false,
		nil)
	subscriber.errorHandler.LogWarning(err, fmt.Sprintf("Error occurred while attempting to consume messages from queue %s", subscriber.config.QueueName))

	forever := make(chan bool)

	go func() {
		for message := range messages {
			distributedMessageType := reflect.TypeOf(distributedMessage)
			if distributedMessageType.Kind() != reflect.Struct {
				subscriber.errorHandler.LogWarning(nil, "Implementation of IDistributedMessage is not a struct")
			}
			iDistributedMessageType := reflect.TypeOf((*models.IDistributedMessage)(nil)).Elem()
			if !distributedMessageType.Implements(iDistributedMessageType) {
				subscriber.errorHandler.LogWarning(nil, "IDistributedMessage implementation does not implement IDistributedMessage interface")
			}

			err := json.Unmarshal(message.Body, &distributedMessage)
			subscriber.errorHandler.LogWarning(err, "An error occurred while trying to load the message body into an instance of IDistributedMessage implementation.")

			handler.HandleMessage(distributedMessage)
			message.Ack(false) //Acknowledge just this message.
		}

		log.Println("Finished processing all messages on queue... ")
		log.Println("[*] Waiting for messages. To exit press CTRL+C")
	}()

	log.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
