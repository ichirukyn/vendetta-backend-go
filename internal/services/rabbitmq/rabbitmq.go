package rabbitmq

import (
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"vendetta/internal/app/config"
	"vendetta/pkg/utils"
)

type Connection = amqp.Connection
type Channel = amqp.Channel
type Queue = amqp.Queue
type Delivery = amqp.Delivery

// Service структура для работы с rabbitmq
type Service struct {
	l *utils.Logger
	c *config.Config

	conn *Connection

	channels map[string]*Channel
	queues   map[string]*Queue

	traceName string
}

// Init подключается к rabbitmq
func (rmq *Service) Init() error {
	url := utils.NewURLConnectionString("amqp", rmq.c.RabbitMQHost, rmq.c.RabbitMQPath, "", rmq.c.RabbitMQUser, rmq.c.RabbitMQPassword)

	conn, err := amqp.Dial(url)
	if err != nil {
		rmq.l.ErrorT(rmq.traceName, "Failed to connect to RabbitMQ", err)
		return err
	}

	rmq.l.InfoT(rmq.traceName, "Successfully connected to RabbitMQ")

	rmq.conn = conn
	return nil
}

// Disconnect отключается от rabbitmq
func (rmq *Service) Disconnect() error {
	if err := rmq.conn.Close(); err != nil {
		rmq.l.ErrorT(rmq.traceName, "Failed to disconnect from RabbitMQ", err)
		return err
	}

	rmq.l.InfoT(rmq.traceName, "Successfully disconnected to RabbitMQ")
	rmq.conn = nil
	return nil
}

// GetClient возвращает клиент rabbitmq RabbitMQConnection
func (rmq *Service) GetClient() *Connection {
	return rmq.conn
}

// CreateChannel создаёт канал с указанным именем
func (rmq *Service) CreateChannel(name string) error {
	if rmq.channels[name] != nil {
		rmq.l.ErrorT(rmq.traceName, "A channel with that name already exists", name)
		return errors.New("the channel already exists")
	}

	ch, err := rmq.conn.Channel()
	if err != nil {
		rmq.l.ErrorT(rmq.traceName, "Failed to create a channel", err)
		return err
	}

	rmq.l.DebugT(rmq.traceName, "The channel has been successfully created")
	rmq.channels[name] = ch
	return nil
}

// CloseChannel закрывает запрашиваемый канал
func (rmq *Service) CloseChannel(name string) error {
	if _, err := rmq.GetChannel(name); err != nil {
		return err
	}

	if err := rmq.channels[name].Close(); err != nil {
		rmq.l.ErrorT(rmq.traceName, "The channel could not be closed", name, err)
		return err
	}

	rmq.l.InfoT(rmq.traceName, "The channel has been successfully closed", name)
	return nil
}

// GetChannel возвращает инстанс запрашиваемого канала (RabbitMQChannel)
func (rmq *Service) GetChannel(name string) (*Channel, error) {
	if rmq.channels[name] == nil {
		rmq.l.ErrorT(rmq.traceName, "A channel with that name was not found", name)
		return nil, errors.New("the channel was not found")
	}

	return rmq.channels[name], nil
}

// CreateQueue создаёт новую очередь в одноименнёном канале
func (rmq *Service) CreateQueue(name string, durable bool) error {
	if _, err := rmq.GetChannel(name); err != nil {
		return err
	}
	if rmq.queues[name] != nil {
		rmq.l.ErrorT(rmq.traceName, "A queue with that name already exist", name)
		return errors.New("the queue already exists")
	}

	q, err := rmq.channels[name].QueueDeclare(
		name,
		durable,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		rmq.l.ErrorT(rmq.traceName, "Failed to create a queue", err)
		return err
	}

	rmq.l.DebugT(rmq.traceName, "The queue has been successfully created")
	rmq.queues[name] = &q
	return nil
}

// GetQueue возвращает инстанс запрашиваемой очереди (RabbitMQQueue)
func (rmq *Service) GetQueue(name string) (*Queue, error) {
	if rmq.queues[name] == nil {
		rmq.l.ErrorT(rmq.traceName, "A queue with this name was not found", name)
		return nil, errors.New("the queue was not found")
	}

	return rmq.queues[name], nil
}

// PublishJSON добавляет в очередь JSON сообщение
func (rmq *Service) PublishJSON(name, exchange string, data interface{}) error {
	if _, err := rmq.GetQueue(name); err != nil {
		return err
	}
	if _, err := rmq.GetChannel(name); err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		rmq.l.ErrorT(rmq.traceName, "Cannot marshal data", name, err)
		return err
	}

	if err := rmq.channels[name].Publish(
		exchange,
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		}); err != nil {
		rmq.l.ErrorT(rmq.traceName, "Failed to publish a message", err)
		return err
	}

	rmq.l.InfoT(rmq.traceName, "The JSON message was successfully published", name)
	return nil
}

// PublishString добавляет в очередь строковое сообщение
func (rmq *Service) PublishString(name, exchange string, data string) error {
	if _, err := rmq.GetQueue(name); err != nil {
		return err
	}
	if _, err := rmq.GetChannel(name); err != nil {
		return err
	}

	if err := rmq.channels[name].Publish(
		exchange,
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		}); err != nil {
		rmq.l.ErrorT(rmq.traceName, "Failed to publish a message", err)
		return err
	}

	rmq.l.InfoT(rmq.traceName, "The message was successfully published", name)
	return nil
}

// AddConsumer добавляет обработчик сообщений из очереди
func (rmq *Service) AddConsumer(qName string, taskLimitCount int, consumer func(d Delivery)) error {
	if _, err := rmq.GetQueue(qName); err != nil {
		return err
	}
	if _, err := rmq.GetChannel(qName); err != nil {
		return err
	}

	msgs, err := rmq.channels[qName].Consume(
		rmq.queues[qName].Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		rmq.l.ErrorT(rmq.traceName, "Error consuming message", qName, err)
		return err
	}

	taskLimiter := make(chan struct{}, taskLimitCount)

	go func() {
		for d := range msgs {
			taskLimiter <- struct{}{}

			go func(d Delivery) {
				defer func() { <-taskLimiter }()
				consumer(d)
			}(d)
		}
	}()

	rmq.l.DebugT(rmq.traceName, "The consumer has been successfully added")
	return nil
}

// NewRabbitMQService возвращает инстанс RabbitMQService
func NewRabbitMQService(l *utils.Logger, c *config.Config) *Service {
	return &Service{
		l: l,
		c: c,

		channels: make(map[string]*Channel),
		queues:   make(map[string]*Queue),

		traceName: "[RabbitMQService]",
	}
}
