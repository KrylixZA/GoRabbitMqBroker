package main

import (
	"time"

	"github.com/KrylixZA/GoRabbitMqBroker/broker"
	"github.com/KrylixZA/GoRabbitMqBroker/enums"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
)

func main() {
	publisherConfig := models.Config{
		Username:     "admin",
		Password:     "admin",
		RabbitMqHost: "localhost",
		VirtualHost:  "/",
		PublisherConfig: &models.PublisherConfig{
			ExchangeName:       "testExchange",
			BindingType:        enums.Topic,
			Durable:            true,
			MandatoryQueueBind: true,
		},
	}

	broker := broker.NewMessagePublisher(publisherConfig)
	defer broker.CloseChannel()
	defer broker.CloseConnection()

	testDataPayload := basicDistributedMessage{
		Data: "My test message",
	}

	broker.Publish("myTestRoutingKey", testDataPayload)
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
