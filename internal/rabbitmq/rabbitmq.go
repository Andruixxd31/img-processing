package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Connect() error
}

type RabbitMQ struct {
	Conn amqp.Connection
}

func (rmq *RabbitMQ) Connect() {
	fmt.Println("Connect method")
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
