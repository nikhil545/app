package database

import (
	"fmt"
	"github.com/streadway/amqp"
)

func StartMessageBroker() *amqp.Connection{
	connStr := fmt.Sprintf("amqp://%s:%s@%s:5672", "guest", "guest","localhost")
	conn, err := amqp.Dial(connStr)
	if err != nil {
		panic(err)
	}

	return conn
}