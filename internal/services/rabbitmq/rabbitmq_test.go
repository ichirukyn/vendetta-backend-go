package rabbitmq_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"vendetta/internal/app/config"
	"vendetta/internal/services/rabbitmq"
	"vendetta/pkg/utils"
)

var Config *config.Config
var Logger *utils.Logger

func TestMain(m *testing.M) {
	dir, _ := os.Getwd()

	Config = config.NewConfig(utils.GetEnvFilePath(dir, ".env.test"))
	Logger = utils.NewDefaultLogger(Config.AppMode)

	os.Exit(m.Run())
}

func TestService_Init(t *testing.T) {
	rmq := rabbitmq.NewRabbitMQService(Logger, Config)
	if err := rmq.Init(); err != nil {
		t.Fatal(err)
	}

	defer func(rmq *rabbitmq.Service) {
		err := rmq.Disconnect()
		if err != nil {
			t.Fatal(err)
		}
	}(rmq)

	assert.NotNil(t, rmq.GetClient())

	err := rmq.CreateChannel("queue")
	assert.NoError(t, err)

	channel, err := rmq.GetChannel("queue")
	assert.NoError(t, err)
	assert.NotNil(t, channel)

	err = rmq.CreateQueue("queue", false)
	assert.NoError(t, err)

	queue, err := rmq.GetQueue("queue")
	assert.NoError(t, err)
	assert.NotNil(t, queue)
}

func TestService_Init2(t *testing.T) {
	rmq := rabbitmq.NewRabbitMQService(Logger, Config)

	assert.Nil(t, rmq.GetClient())
}
