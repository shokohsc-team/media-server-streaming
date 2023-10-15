package queue

import (
	"os"
    "log"

    amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
    	log.Panicf("%s: %s", msg, err)
    }
}

func ConnectAMQP() (*amqp.Connection, *amqp.Channel) {
	env := os.Getenv("AMQP_URL")
	conn, err := amqp.Dial(env)
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()
	return conn, ch
}

func CloseAMQP(conn *amqp.Connection, channel *amqp.Channel) {
    defer conn.Close()    //rabbit mq close
    defer channel.Close() //rabbit mq channel close
}

