package main

import (
	"log"

	"github.com/KrylixZA/GoRabbitMqBroker/bindingType"
	"github.com/KrylixZA/GoRabbitMqBroker/broker"
	"github.com/KrylixZA/GoRabbitMqBroker/logs"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
	"github.com/KrylixZA/GoRabbitMqBroker/processing"
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
			BindingType:     bindingType.Topic,
			RoutingKey:      "*",
			PrefetchCount:   1,
			StrictQueueName: true,
			Durable:         true,
			AutoDeleteQueue: false,
			RequeueOnNack:   true,
		},
	}

	broker := broker.NewMessageSubscriber(subscriberConfig, logs.Logger{})
	defer broker.Close()

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
