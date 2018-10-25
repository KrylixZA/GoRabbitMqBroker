package processing

import "github.com/KrylixZA/GoRabbitMqBroker/models"

//IMessageHandler describes a contract that all subscribers of a RabbitMQ queue must implement.
//The subscribers must pass a struct which implements IDistributedMessage.
//The message handler will be called when a message is ready to be consumed.
//The message handler will be called in an asynchronous manner.
type IMessageHandler interface {
	HandleMessage(disributedMessage models.DistributedMessage) error
}
