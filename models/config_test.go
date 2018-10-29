package models

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrylixZA/GoRabbitMqBroker/bindingType"
)

func TestValidate_GivenValidConfig_ShouldReturnNil(t *testing.T) {
	// Arrange
	config := Config{
		Username:     "test",
		Password:     "test",
		RabbitMqHost: "localhost",
		VirtualHost:  "/",
		SubscriberConfig: &SubscriberConfig{
			QueueName:       "test",
			ExchangeName:    "test",
			BindingType:     bindingtype.Topic,
			RoutingKey:      "test.*",
			PrefetchCount:   100,
			StrictQueueName: true,
			Durable:         true,
			AutoDeleteQueue: false,
			RequeueOnNack:   true,
		},
		PublisherConfig: &PublisherConfig{
			ExchangeName:       "test",
			BindingType:        bindingtype.Fanout,
			Durable:            true,
			MandatoryQueueBind: false,
		},
	}

	// Act
	err := config.Validate()

	// Assert
	assert.Nil(t, err)
}

func TestValidate_GivenBadUsername_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	config := Config{
		Username:     "",
		Password:     "test",
		RabbitMqHost: "localhost",
		VirtualHost:  "/",
		SubscriberConfig: &SubscriberConfig{
			QueueName:       "test",
			ExchangeName:    "test",
			BindingType:     bindingtype.Topic,
			RoutingKey:      "test.*",
			PrefetchCount:   100,
			StrictQueueName: true,
			Durable:         true,
			AutoDeleteQueue: false,
			RequeueOnNack:   true,
		},
		PublisherConfig: &PublisherConfig{
			ExchangeName:       "test",
			BindingType:        bindingtype.Fanout,
			Durable:            true,
			MandatoryQueueBind: false,
		},
	}
	expectedError := errors.New("username is empty string")

	// Act
	err := config.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidate_GivenBadPassword_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	config := Config{
		Username:     "test",
		Password:     "",
		RabbitMqHost: "localhost",
		VirtualHost:  "/",
		SubscriberConfig: &SubscriberConfig{
			QueueName:       "test",
			ExchangeName:    "test",
			BindingType:     bindingtype.Topic,
			RoutingKey:      "test.*",
			PrefetchCount:   100,
			StrictQueueName: true,
			Durable:         true,
			AutoDeleteQueue: false,
			RequeueOnNack:   true,
		},
		PublisherConfig: &PublisherConfig{
			ExchangeName:       "test",
			BindingType:        bindingtype.Fanout,
			Durable:            true,
			MandatoryQueueBind: false,
		},
	}
	expectedError := errors.New("password is empty string")

	// Act
	err := config.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidate_GivenBadRabbitMqHost_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	config := Config{
		Username:     "test",
		Password:     "test",
		RabbitMqHost: "",
		VirtualHost:  "/",
		SubscriberConfig: &SubscriberConfig{
			QueueName:       "test",
			ExchangeName:    "test",
			BindingType:     bindingtype.Topic,
			RoutingKey:      "test.*",
			PrefetchCount:   100,
			StrictQueueName: true,
			Durable:         true,
			AutoDeleteQueue: false,
			RequeueOnNack:   true,
		},
		PublisherConfig: &PublisherConfig{
			ExchangeName:       "test",
			BindingType:        bindingtype.Fanout,
			Durable:            true,
			MandatoryQueueBind: false,
		},
	}
	expectedError := errors.New("host is empty string")

	// Act
	err := config.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidate_GivenBadVirtualHost_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	config := Config{
		Username:     "test",
		Password:     "test",
		RabbitMqHost: "test",
		VirtualHost:  "",
		SubscriberConfig: &SubscriberConfig{
			QueueName:       "test",
			ExchangeName:    "test",
			BindingType:     bindingtype.Topic,
			RoutingKey:      "test.*",
			PrefetchCount:   100,
			StrictQueueName: true,
			Durable:         true,
			AutoDeleteQueue: false,
			RequeueOnNack:   true,
		},
		PublisherConfig: &PublisherConfig{
			ExchangeName:       "test",
			BindingType:        bindingtype.Fanout,
			Durable:            true,
			MandatoryQueueBind: false,
		},
	}
	expectedError := errors.New("vhost is empty string")

	// Act
	err := config.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidate_GivenNilSubscriberConfigAndNilPublisherConfig_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	config := Config{
		Username:     "test",
		Password:     "test",
		RabbitMqHost: "test",
		VirtualHost:  "/",
	}
	expectedError := errors.New("subscriberConfig and publisherConfig are missing. A consumer of the RabbitMQ broker must be a producer, or a consumer, or both")

	// Act
	err := config.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidateSubscriberConfig_GivenValidSubscriberConfig_ShouldReturnNil(t *testing.T) {
	// Arrange
	subscriberConfig := SubscriberConfig{
		QueueName:       "test",
		ExchangeName:    "test",
		BindingType:     bindingtype.Topic,
		RoutingKey:      "test.*",
		PrefetchCount:   100,
		StrictQueueName: true,
		Durable:         true,
		AutoDeleteQueue: false,
		RequeueOnNack:   true,
	}

	// Act
	err := subscriberConfig.Validate()

	// Assert
	assert.Nil(t, err)
}

func TestValidateSubscriberConfig_GivenStrictQueueNameAndBadQueueName_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	subscriberConfig := SubscriberConfig{
		QueueName:       "",
		ExchangeName:    "test",
		BindingType:     bindingtype.Topic,
		RoutingKey:      "test.*",
		PrefetchCount:   100,
		StrictQueueName: true,
		Durable:         true,
		AutoDeleteQueue: false,
		RequeueOnNack:   true,
	}
	expectedError := errors.New("subscriberConfig.strictQueueName is set to true but subscriberConfig.queueName is empty string. If you wish to use auto-generated queue names, set strictQueueName to false")

	// Act
	err := subscriberConfig.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidateSubscriberConfig_GivenBadExchangeName_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	subscriberConfig := SubscriberConfig{
		QueueName:       "test",
		ExchangeName:    "",
		BindingType:     bindingtype.Topic,
		RoutingKey:      "test.*",
		PrefetchCount:   100,
		StrictQueueName: true,
		Durable:         true,
		AutoDeleteQueue: false,
		RequeueOnNack:   true,
	}
	expectedError := errors.New("subscriberConfig.exchangeName is empty string. Although RabbitMQ allows for auto-generating exchange names, it becomes complex to manage when binding queues. As such, we force an exchangeName to be supplied in the config")

	// Act
	err := subscriberConfig.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidateSubscriberConfig_GivenBadBindingType_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	subscriberConfig := SubscriberConfig{
		QueueName:       "test",
		ExchangeName:    "test",
		BindingType:     -1,
		RoutingKey:      "test.*",
		PrefetchCount:   100,
		StrictQueueName: true,
		Durable:         true,
		AutoDeleteQueue: false,
		RequeueOnNack:   true,
	}
	expectedError := errors.New("subscriberConfig.bindingType is out of range. Acceptable options are 0 = Fanout, 1 = Direct, 2 = Topic")

	// Act
	err := subscriberConfig.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidateSubscriberConfig_GivenDirectBindingAndEmptyRoutingKey_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	subscriberConfig := SubscriberConfig{
		QueueName:       "test",
		ExchangeName:    "test",
		BindingType:     bindingtype.Direct,
		RoutingKey:      "",
		PrefetchCount:   100,
		StrictQueueName: true,
		Durable:         true,
		AutoDeleteQueue: false,
		RequeueOnNack:   true,
	}
	expectedError := errors.New("subscriberConfig.routingKey is empty string. Cannot use an empty routing key to bind a queue to an exchange when using Direct or Topic based routing")

	// Act
	err := subscriberConfig.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidateSubscriberConfig_GivenTopicBindingAndEmptyRoutingKey_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	subscriberConfig := SubscriberConfig{
		QueueName:       "test",
		ExchangeName:    "test",
		BindingType:     bindingtype.Topic,
		RoutingKey:      "",
		PrefetchCount:   100,
		StrictQueueName: true,
		Durable:         true,
		AutoDeleteQueue: false,
		RequeueOnNack:   true,
	}
	expectedError := errors.New("subscriberConfig.routingKey is empty string. Cannot use an empty routing key to bind a queue to an exchange when using Direct or Topic based routing")

	// Act
	err := subscriberConfig.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidateSubscriberConfig_GivenNegativePrefetchCount_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	subscriberConfig := SubscriberConfig{
		QueueName:       "test",
		ExchangeName:    "test",
		BindingType:     bindingtype.Topic,
		RoutingKey:      "test.*",
		PrefetchCount:   -1,
		StrictQueueName: true,
		Durable:         true,
		AutoDeleteQueue: false,
		RequeueOnNack:   true,
	}
	expectedError := errors.New("subscriberConfig.prefetchCount cannot be less than zero")

	// Act
	err := subscriberConfig.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidatePublisherConfig_GivenValidPublisherConfig_ShouldReturnNil(t *testing.T) {
	// Arrange
	publisherConfig := PublisherConfig{
		ExchangeName:       "test",
		BindingType:        bindingtype.Fanout,
		Durable:            true,
		MandatoryQueueBind: false,
	}

	// Act
	err := publisherConfig.Validate()

	// Assert
	assert.Nil(t, err)
}

func TestValidatePublisherConfig_GivenBadExchangeName_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	publisherConfig := PublisherConfig{
		ExchangeName:       "",
		BindingType:        bindingtype.Fanout,
		Durable:            true,
		MandatoryQueueBind: false,
	}
	expectedError := errors.New("publisherConfig.exchangeName is empty string. Although RabbitMQ allows for auto-generating exchange names, it becomes complex to manage when binding queues. As such, we force an exchangeName to be supplied in the config")

	// Act
	err := publisherConfig.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}

func TestValidatePublisherConfig_GivenBadBindingType_ShouldReturnExpectedError(t *testing.T) {
	// Arrange
	publisherConfig := PublisherConfig{
		ExchangeName:       "test",
		BindingType:        3,
		Durable:            true,
		MandatoryQueueBind: false,
	}
	expectedError := errors.New("publisherConfig.bindingType is out of range. Acceptable options are 0 = Fanout, 1 = Direct, 2 = Topic")

	// Act
	err := publisherConfig.Validate()

	// Assert
	assert.Equal(t, expectedError, err)
}
