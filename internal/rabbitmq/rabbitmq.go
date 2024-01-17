package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Connect()
	DeclareQueue(queueName string)
	DeclareExchange(exchangeName string)
	SetBinding(queueName, routingKey, exchangeName string)
	SetQos(prefetchCount int)
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
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		failOnError(err, "Could not declare queue")
	}
}

func (rmq *RabbitMQ) DeclareExchange(exchangeName string) {
	err := rmq.Channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		failOnError(err, "Could not declare exchange")
	}
}

func (rmq *RabbitMQ) SetBinding(queueName, routingKey, exchangeName string) {
	err := rmq.Channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)

	if err != nil {
		failOnError(err, "Could not declare exchange")
	}
}

func (rmq *RabbitMQ) SetQos(prefetchCount int) {
	err := rmq.Channel.Qos(
		1,
		0,
		false,
	)

	if err != nil {
		failOnError(err, "Could not set Qos")
	}
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
