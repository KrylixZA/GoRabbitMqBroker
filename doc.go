// Package gorabbitmqbroker exposes an abstracted and simplified implementation for interacting with a RabbitMQ node or cluster.
// At a very high level, a user will create/read a configuration which defines whether they intend to be a publisher or a subscriber
// The user will pass this configuration to the message broker which will handle the details of connecting to RabbitMQ and defining the queues and exchanges.
// After this, the user provides the necessary interfaces to the broker and it will handle the details of subscribing and reading messages or publishing messages as per the config.

// Known issues can be found [here](https://github.com/KrylixZA/GoRabbitMqBroker/issues)

// This code is licensed under an MIT license.

// Authors: Simon Headley (KrylixZA)

package gorabbitmqbroker
