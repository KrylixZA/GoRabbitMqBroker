//Package bindingtype exposes an enumerable that represents how a queue is bound to an exchange.
//The purpose of this package is to simply the user experience of the user when setting up their configuration for connection to RabbitMQ.
//Known issues can be found on GitHub (https://github.com/KrylixZA/GoRabbitMqBroker/issues).
//This code is licensed under an MIT license.
//Authors: Simon Headley (KrylixZA).
package bindingtype

//BindingType defines the type of binding a queue will use when binding itself to a queue.
//		An exchange can be bound to any number of queues, but all must use the same binding type.
//		In the case of direct and topic bindings, all bindings are white-listed.
//		This means that for multiple routes to lead to a queue, they must either use topic based or all be verbosely specified.
//Default bindingType is fanout
type BindingType int

const (
	//Fanout requires no routing keys. All messages published to an exchange will be published to the queue.
	Fanout BindingType = iota

	//Direct requires an exact match of the routing key. Only messages published to the exchange with that exact routing key will be routed to the queue.
	Direct

	//Topic is similar to direct, however it accepts wildcard routes.
	//		This means that a subset of messages that follow a particular routing key pattern will all be routed to the queue.
	//		If no wildcard character is supplied, the binding will behave the same as direct.
	//		Traditional wildcard characters are "*" and "#". More can be read here: https://www.rabbitmq.com/tutorials/tutorial-five-go.html
	Topic
)

func (bindingType BindingType) String() string {
	return [...]string{"fanout", "direct", "topic"}[bindingType]
}
