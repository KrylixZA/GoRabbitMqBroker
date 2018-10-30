# GoRabbitMqBroker [![GoDoc](https://godoc.org/github.com/KrylixZA/GoRabbitMqBroker?status.svg)](https://godoc.org/github.com/KrylixZA/GoRabbitMqBroker) [GitHub Pages](https://krylixza.github.io/GoRabbitMqBroker/)
A library of code which abstracts and encapsulates the details of connecting and configuration an AMQP connection a RabbitMQ Broker while still providing the majority of the flexibility gained by referring to the Go AMQP library directly.

## To install
```bash
$ go get github.com/KrylixZA/GoRabbitMqBroker
```

### Setting up RabbitMQ
1. Install RabbitMQ server (either yum, homebrew, etc...). If installed through yum, run by executing:
```bash
$ rabbitmq-server
```
If installed through homebrew, run by executing:
```bash
$ brew services start rabbitmq
```
2. Browse to http://localhost:15672 (this the RabbitMQ management portal).
3. Login using "guest", "guest".
4. Create a user "admin" with password "admin".
5. Give the "admin" user full permissions to the Virtual Host "/".

### Using this package
* Before rambo-ing your way through this code, please take a minute to check out the [examples](https://github.com/KrylixZA/GoRabbitMqBroker/tree/master/examples) and read all the comments on all public-facing endpoints.

1. All your code will communicate with RabbitMQ through [messageBroker.go](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/broker/messageBroker.go).
2. You will need to make use of the configuration defined [here](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/models/config.go).
3. You will, likely, also need to make use of the Binding Type enumeration defined [here](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/bindingType/bindingTypes.go).
4. Publishers need only interact with the [NewMessagePublisher definition](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/broker/messageBroker.go#L70).
5. Subscribers will need to interact with [NewMessageSubscriber definition](hhttps://github.com/KrylixZA/GoRabbitMqBroker/blob/master/broker/messageBroker.go#L41) and the [IMessageHandler interface](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/processing/messageHandler.go#L9).
6. Publishers and subscribes will need to interact with [NewMessagePublisherSubscriber definition](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/broker/messageBroker.go#L94) and the [IMessageHandler interface](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/processing/messageHandler.go#L9).
7. All messages that flow through RabbitMQ via the [messageBroker.go](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/broker/messageBroker.go) are an implementation of the [IDistributedMessage interface](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/models/distributedMessage.go#L24).
8. All subscribers receive a concrete implementation of [IDistributedMessage](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/models/distributedMessage.go#L24) in the form of [DistributedMessage](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/models/distributedMessage.go#L24).
9. All publishers must publish a struct which implements [IDistributedMessage](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/models/distributedMessage.go#L24).

### Examples
1. An example of a basic publisher can be found [here](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/examples/basicPublisher.go). To run this:
```bash
$ cd ${GOPATH}/src/github.com/KrylixZA/GoRabbitMqBroker/examples
$ go run basicPublisher.go
```
2. An example of a basic subscriber can be found [here](https://github.com/KrylixZA/GoRabbitMqBroker/blob/master/examples/basicSubscriber.go). To run this:
```bash
$ cd ${GOPATH}/src/github.com/KrylixZA/GoRabbitMqBroker/examples
$ go run basicSubscriber.go
```