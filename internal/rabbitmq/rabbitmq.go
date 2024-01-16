package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Connect() error
	DeclareQueue(queueName string) error
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func failOnError(err error, msg string) {
	log.Panicf("%s: %s", msg, err)
}

func (rmq *RabbitMQ) Connect() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		failOnError(err, "Could not establish connection with rabbitmq")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Could not open channel")
	}
	defer ch.Close()

	rmq.Channel = ch
}

func (rmq *RabbitMQ) DeclareQueue(queueName string) {
	_, err := rmq.Channel.QueueDeclare(
		"first",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		failOnError(err, "Could not declare queue")
	}
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
