package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/KrylixZA/GoRabbitMqBroker/bindingType"
	"github.com/KrylixZA/GoRabbitMqBroker/broker"
	"github.com/KrylixZA/GoRabbitMqBroker/logs"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
	uuid "github.com/satori/go.uuid"
)

func main() {
	numOfMsgsToPublish := flag.Int("n", 1, "Defines the number of messages to publish to RabbitMQ.")
	flag.Parse()
	log.Printf("Publishing %d messages...", *numOfMsgsToPublish)

	publisherConfig := models.Config{
		Username:     "admin",
		Password:     "admin",
		RabbitMqHost: "localhost",
		VirtualHost:  "/",
		PublisherConfig: &models.PublisherConfig{
			ExchangeName:       "amq.topic",
			BindingType:        bindingType.Topic,
			Durable:            true,
			MandatoryQueueBind: false,
		},
	}

	broker := broker.NewMessagePublisher(publisherConfig, logs.Logger{})
	defer broker.Close()

	for i := 0; i < *numOfMsgsToPublish; i++ {
		testDataPayload := basicDistributedMessage{
			Data: fmt.Sprintf("[%d] My test message", i),
		}

		broker.Publish("myTestRoutingKey", testDataPayload)
	}
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
