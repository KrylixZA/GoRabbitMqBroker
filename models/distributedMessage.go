package models

import "time"

//IDistributedMessage defines a contract that all messages which are published or consumed from RabbitMQ will adhere to.
//GetData is a function that returns the payload of the message.
//		In the case of a message being published, this will be the struct that represents the body to be published.
//		In the case of a message being consumed, this will be the data payload that was used to create the body.
//GetTimestamp is a function that returns a timestamp of significance that gives context to the message.
//		GetTimestamp is open to interpretation. You are free to design the implementation as you would like, or you can leave it unimplemented.
//		If left unimplemented, be sure not to make use of it in your implementation of IMessageHandler
//GetCorrelationId is a function that returns the correlationId of the message.
//		CorrelationId is designed for Remote Procedure Calls, it's primary purpose is correlate the message to the request/response it belongs.
//		CorrelationId can be used when sharing messages between systems, which allow remote systems to return a message to originated with a state, or:
//		CorrelationId can be used when a request is distributed and passed through the RabbitMQ stream that you define and you wish to keep track of all the messages that belong to the original request, correlationId should be used.
//		CorrelationId is a string so it can accept any kind of unique identifier. A strong argument can be made for using GUIDs/UUIDs in this scenario.
//		You can read up more about CorrelationId's here: https://www.rabbitmq.com/tutorials/tutorial-six-go.html
type IDistributedMessage interface {
	GetData() interface{}
	GetTimestamp() time.Time
	GetCorrelationId() string
}
