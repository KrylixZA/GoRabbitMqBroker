package broker

import (
	"encoding/json"
	"fmt"

	"github.com/KrylixZA/GoRabbitMqBroker/logs"
	"github.com/KrylixZA/GoRabbitMqBroker/models"
	"github.com/streadway/amqp"
)

type messagePublisher struct {
	config  models.PublisherConfig
	channel *amqp.Channel
	logger  logs.ILogger
}

func newMessagePublisher(config models.PublisherConfig, channel *amqp.Channel, logger logs.ILogger) *messagePublisher {
	publisher := messagePublisher{
		config:  config,
		channel: channel,
		logger:  logger,
	}

	//Declare the exchange
	err := channel.ExchangeDeclare(
		config.ExchangeName,
		config.BindingType.String(),
		config.Durable,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}

	return &publisher
}

func (publisher *messagePublisher) publish(routingKey string, distributedMessage models.IDistributedMessage) error {
	distributedMessageJSONPayload, err := json.Marshal(distributedMessage.GetData())
	if err != nil {
		publisher.logger.LogWarning(fmt.Sprintf("Error occurred while creating JSON payload from distributedMessage %s\n\n%s",
			distributedMessage,
			err))
	}

	publishParams := amqp.Publishing{
		DeliveryMode:  amqp.Persistent,
		ContentType:   "text/json",
		CorrelationId: distributedMessage.GetCorrelationId(),
		MessageId:     distributedMessage.GetMessageId(),
		Timestamp:     distributedMessage.GetTimestamp(),
		Body:          distributedMessageJSONPayload,
	}
	err = publisher.channel.Publish(
		publisher.config.ExchangeName,
		routingKey,
		publisher.config.MandatoryQueueBind,
		false,
		publishParams)

	if err != nil {
		publisher.logger.LogWarning(fmt.Sprintf("Error occurred while publishing args=%+v to exchange=%s with routing key=%s\n\n%s",
			publishParams,
			publisher.config.ExchangeName,
			routingKey,
			err))
	}

	return nil
}
