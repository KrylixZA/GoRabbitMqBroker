package main

import (
	"time"

	"github.com/KrylixZA/GoRabbitMqBroker/broker"
	"github.com/KrylixZA/GoRabbitMqBroker/enums"
	"github.com/KrylixZA/GoRabbitMqBroker/logs"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
	uuid "github.com/satori/go.uuid"
)

func main() {
	publisherConfig := models.Config{
		Username:     "admin",
		Password:     "admin",
		RabbitMqHost: "localhost",
		VirtualHost:  "/",
		PublisherConfig: &models.PublisherConfig{
			ExchangeName:       "amq.topic",
			BindingType:        enums.Topic,
			Durable:            true,
			MandatoryQueueBind: false,
		},
	}

	broker := broker.NewMessagePublisher(publisherConfig, logs.Logger{})
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

func (distributedMessage basicDistributedMessage) GetMessageId() string {
	uuId, _ := uuid.NewV4()

	return uuId.String()
}

func (distributedMessage basicDistributedMessage) GetCorrelationId() string {
	uuId, _ := uuid.NewV4()

	return uuId.String()
}
