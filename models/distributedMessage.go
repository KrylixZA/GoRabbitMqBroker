package models

import "time"

//IDistributedMessage defines a contract that all messages which are published or consumed from RabbitMQ will adhere to.
//GetData is a function that returns the payload of the message.
//		In the case of a message being published, this will be the struct that represents the body to be published.
//		In the case of a message being consumed, this will be the data payload that was used to create the body.
//GetTimestamp is a function that returns a timestamp of significance that gives context to the message.
//		GetTimestamp is open to interpretation. You are free to design the implementation as you would like.
//		If you have no use for a timestamp, simply return time.Now()
//		If this function is left unimplemented, publishing messages will fail.
//GetMessageId is a function that returns a unique identifier for this particular message.
//		Typically MessageId is an idempotent identifier that can be used to refer back to a single request.
//		MessageId holds no significance aside from that it must be unique per message (not necessarily per request).
//		If a message is split apart and distributed into multiple messages, each distributed message must have a unique MessageId.
//		MessageId is a string so it can accept any kind of unique identifier. A strong argument can be made for using GUIDs/UUIDs in this scenario.
//GetCorrelationId is a function that returns the correlationId of the message.
//		CorrelationId is designed for Remote Procedure Calls, it's primary purpose is correlate the message to the request/response it belongs.
//		CorrelationId can be used when sharing messages between systems, which allow remote systems to return a message to originated with a state, or:
//		CorrelationId can be used when a request is distributed and passed through the RabbitMQ stream that you define and you wish to keep track of all the messages that belong to the original request, correlationId should be used.
//		CorrelationId is a string so it can accept any kind of unique identifier. A strong argument can be made for using GUIDs/UUIDs in this scenario.
//		You can read up more about CorrelationId's here: https://www.rabbitmq.com/tutorials/tutorial-six-go.html
type IDistributedMessage interface {
	GetData() interface{}
	GetTimestamp() time.Time
	GetMessageId() string
	GetCorrelationId() string
}

//DistributedMessage represents a Go struct that implements the basic requirements of the IDistributedMessage interface.
//Data is of type interface{}, meaning it can contain anything as the data payload.
//Timestamp is of time.Time. This significance of this time is only within the context of it's use.
//CorrelationId is any string uniquely identifying the message to its source.
type DistributedMessage struct {
	Data          interface{} `json:"data"`
	Timestamp     time.Time   `json:"timestamp"`
	MessageId     string      `json:"messageId"`
	CorrelationId string      `json:"correlationId"`
}

//GetData is a raw implementation of the GetData() function defined in IDistributedMessage above.
func (distributedMessage DistributedMessage) GetData() interface{} {
	return distributedMessage.Data
}

//GetTimestamp is a raw implementation of the GetTimestamp() function defined in IDistributedMessage above.
func (distributedMessage DistributedMessage) GetTimestamp() time.Time {
	return distributedMessage.Timestamp
}

//GetMessageId is a raw implementation of the GetMessageId() function defined in IDistributedMessage above.
func (distributedMessage DistributedMessage) GetMessageId() string {
	return distributedMessage.MessageId
}

//GetCorrelationId is a raw implementation of the GetCorrelationId() function defined in IDistributedMessage above.
func (distributedMessage DistributedMessage) GetCorrelationId() string {
	return distributedMessage.CorrelationId
}
