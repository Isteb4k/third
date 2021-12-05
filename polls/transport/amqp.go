package transport

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Client interface {
	Consume()
	Stop()
}

type amqper struct {
	channels map[string]*amqp.Channel
}

func New() Client {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("users", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

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

func (a *amqper) Consume() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("users", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			addTask := &struct {
				Message string
			}{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("Result is %s", addTask.Message)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

		}
	}()

	// Остановка для завершения программы
	<-stopChan

}
