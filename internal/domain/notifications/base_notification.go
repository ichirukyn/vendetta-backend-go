package notifications

import (
	"vendetta/internal/services/rabbitmq"
)

type BaseNotification interface {
	Consumer(d rabbitmq.Delivery)
}
