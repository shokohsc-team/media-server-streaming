package queue

import (
	"context"
	"fmt"
	"log"
	"time"
	"encoding/json"
	"strconv"

	"netflix/database"
	"netflix/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(v models.Video) {
	conn, ch := ConnectAMQP()
	defer CloseAMQP(conn, ch)

	q, err := ch.QueueDeclare(
		"dev-netflix", // name
		true,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := fmt.Sprintf("{\"Path\": \"%s\", \"LibraryID\": \"%d\"}", v.Path, v.LibraryID)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
	})

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func Consume() {
	conn, ch := ConnectAMQP()
	defer CloseAMQP(conn, ch)

	q, err := ch.QueueDeclare(
		"dev-netflix", // name
		true,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to declare a queue")

	var forever chan struct{}

	go func() {
		type video struct {
			Path string
			LibraryID  string
		}

		db := database.DB
		for msg := range msgs {
			log.Printf("1.Received a message: %s", msg.Body)
			vs := new(video)
			err := json.Unmarshal(msg.Body, &vs)
			if err != nil {
				log.Printf("Cannot json.Unmarshal: %s", err)
			}
			libraryID, err := strconv.ParseUint(vs.LibraryID, 10, 64)
			v := models.Video{
				Path:      vs.Path,
				LibraryID: libraryID,
			}
			err = db.Omit("CollectionID").Create(&v).Error
			if err != nil {
				log.Printf("Cannot create: %s", err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}