package main

import (
	"log"

	"github.com/KrylixZA/GoRabbitMqBroker/enums"
	"github.com/KrylixZA/GoRabbitMqBroker/errors"

	"github.com/KrylixZA/GoRabbitMqBroker/processing"

	"github.com/KrylixZA/GoRabbitMqBroker/broker"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
)

func main() {
	subscriberConfig := models.Config{
		Username:     "admin",
		Password:     "admin",
		RabbitMqHost: "localhost",
		VirtualHost:  "/",
		SubscriberConfig: &models.SubscriberConfig{
			QueueName:       "myTestQueue",
			ExchangeName:    "amq.topic",
			BindingType:     enums.Topic,
			RoutingKey:      "*",
			PrefetchCount:   1,
			StrictQueueName: true,
			Durable:         true,
			AutoDeleteQueue: false,
			RequeueOnNack:   true,
		},
	}

	broker := broker.NewMessageSubscriber(subscriberConfig, errors.LogErrorHandler{})
	defer broker.CloseChannel()
	defer broker.CloseConnection()

	var subscriber processing.IMessageHandler
	subscriber = basicSubscriber{}
	broker.Subscribe(subscriber)
}

type basicSubscriber struct {
}

func (subscriber basicSubscriber) HandleMessage(disributedMessage models.DistributedMessage) error {
	log.Printf("%+v", disributedMessage)

	return nil
}
