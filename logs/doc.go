// Package logs exposes a simple interface and an example implementation of the interface which the RabbitMQ broker library uses to report events.
// The implementation of the ILogger interface is up the end user. The example given can be used if desired (as in the examples).
// If the user wishes to supply their own implementation, they must simply provide a logger that implements the interface.
// This flexibility allows the user to extend on the logger by potentially wrapping the logger with metrics and audting, etc.

// Known issues can be found [here](https://github.com/KrylixZA/GoRabbitMqBroker/issues)

// This code is licensed under an MIT license.

// Authors: Simon Headley (KrylixZA)

package logs
