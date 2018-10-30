// Package processing exposes an interface IMessageHandler.
// The IMessageHandler interface enforces the contract that any subscribers must adhere to.

// Known issues can be found [here](https://github.com/KrylixZA/GoRabbitMqBroker/issues)

// This code is licensed under an MIT license.

// Authors: Simon Headley (KrylixZA)

package processing

import "github.com/KrylixZA/GoRabbitMqBroker/models"

//IMessageHandler describes a contract that all subscribers of a RabbitMQ queue must implement.
//The subscribers must pass a struct which implements IDistributedMessage.
//The message handler will be called when a message is ready to be consumed.
//The message handler will be called in an asynchronous manner.
type IMessageHandler interface {
	HandleMessage(disributedMessage models.DistributedMessage) error
}
