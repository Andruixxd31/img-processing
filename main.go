package main

import (
	"fmt"

	"github.com/andruixxd31/concurrency-basics/img-processing/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func main() {
	Run()
}

func Run() {
	fmt.Println("Starting service")
	rmq := rabbitmq.NewRabbitMQService()

	app := App{
		Rmq: rmq,
	}
	app.Rmq.Connect()
}
