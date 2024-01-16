package main

import (
	"fmt"

	"github.com/andruixxd31/concurrency-basics/img-processing/internal/rabbitmq"
)

type RabbitMq struct {
	rmq *rabbitmq.RabbitMQ
}

func main() {
	fmt.Println("Running rabbitmq")
}
