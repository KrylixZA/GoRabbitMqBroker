package main

import (
	"log"
	"time"

	"github.com/KrylixZA/GoRabbitMqBroker/enums"

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
			QueueName:       "",
			ExchangeName:    "testExchange",
			BindingType:     enums.Topic,
			RoutingKey:      "*",
			PrefetchCount:   1,
			StrictQueueName: false,
			Durable:         false,
			AutoDeleteQueue: false,
		},
	}

	broker := broker.NewMessageSubscriber(subscriberConfig)
	defer broker.CloseChannel()
	defer broker.CloseConnection()

	var subscriber processing.IMessageHandler
	subscriber = basicSubscriber{}

	broker.Subscribe(subscriber, basicDistributedMessage{})
}

type basicSubscriber struct {
}

func (subscriber basicSubscriber) HandleMessage(disributedMessage models.IDistributedMessage) error {
	log.Println(disributedMessage.GetData())

	return nil
}

type basicDistributedMessage struct {
	Data string `json:"data"`
}

func (distributedMessage basicDistributedMessage) GetData() interface{} {
	return distributedMessage.Data
}

func (distributedMessage basicDistributedMessage) GetTimestamp() time.Time {
	return time.Now()
}

func (distributedMessage basicDistributedMessage) GetCorrelationId() string {
	return ""
}
