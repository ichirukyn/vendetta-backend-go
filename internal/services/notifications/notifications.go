package notifications

import (
	"vendetta/internal/services/rabbitmq"
	"vendetta/pkg/utils"
)

// Service структура для работы с уведомлениями
type Service struct {
	l *utils.Logger

	rmq *rabbitmq.Service

	name      string
	taskLimit int

	traceName string
}

// Init инициализирует новый канал и очередь для уведомлений
func (n *Service) Init() error {
	if err := n.rmq.CreateChannel(n.name); err != nil {
		n.l.ErrorT(n.traceName, "cannot create the channel", n.name, err)
		return err
	}

	if err := n.rmq.CreateQueue(n.name, true); err != nil {
		n.l.ErrorT(n.traceName, "cannot create the queue", n.name, err)
		return err
	}

	n.l.DebugT(n.traceName, "Notification service has been successfully initialized", n.name)
	return nil
}

// AddConsume инициализирует новый обработчик уведомлений
func (n *Service) AddConsume(consumer func(delivery rabbitmq.Delivery)) error {
	if err := n.rmq.AddConsumer(n.name, n.taskLimit, consumer); err != nil {
		n.l.ErrorT(n.traceName, "cannot add the consumer", n.name, err)
		return err
	}

	n.l.DebugT(n.traceName, "The consumer has been successfully initialized", n.name)
	return nil
}

// Notify добавляет уведомление в очередь на обработку
func (n *Service) Notify(data interface{}) {
	if err := n.rmq.PublishJSON(n.name, "", data); err != nil {
		n.l.ErrorT(n.traceName, "cannot to notify", err)
	}
}

// NewNotificationsService возвращает инстанс NotificationsService
func NewNotificationsService(l *utils.Logger, rmq *rabbitmq.Service, taskLimit int, name string) *Service {
	return &Service{
		l: l,

		rmq: rmq,

		name:      name,
		taskLimit: taskLimit,

		traceName: "[NotificationService]",
	}
}
