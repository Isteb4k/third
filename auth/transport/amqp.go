package transport

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Client interface {
	Publish()
	Stop()
}

type amqper struct {
	channels map[string]*amqp.Channel
}

func New() Client {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	handleError(err, "Can't connect to AMQP")

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	queue, err := amqpChannel.QueueDeclare(
		"users",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not declare `users` queue")

	channels := make(map[string]*amqp.Channel)
	channels[queue.Name] = amqpChannel

	return &amqper{
		channels: channels,
	}
}

func (a *amqper) Stop() {
	for _, ch := range a.channels {
		err := ch.Close()
		if err != nil {
			log.Println("Failed to close channel:", err)
		}
	}
}

func (a *amqper) Publish() {
	addTask := struct {
		Message string
	}{
		Message: "ping",
	}

	body, err := json.Marshal(addTask)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	queue := "users"

	err = a.channels[queue].Publish("", queue, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
		return
	}

	log.Printf("AddTask: %s", addTask.Message)
}
