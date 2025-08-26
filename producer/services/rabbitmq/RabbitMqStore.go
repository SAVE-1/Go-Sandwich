package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQStore struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var RabbitMQClient *RabbitMQStore

func (r *RabbitMQStore) CloseConnection() {
	r.Channel.Close()
	r.Connection.Close()
}

func NewRabbitMqStore() error {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}

	connection, err := amqp.Dial(url)

	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := connection.Channel()

	if err != nil {
		return fmt.Errorf("failed to open to channel: %w", err)
	}

	RabbitMQClient = &RabbitMQStore{
		Connection: connection,
		Channel:    channel,
	}

	return nil
}

func (r *RabbitMQStore) MakeASandwichRequest(o ObjectRequest) error {
	q, err := r.Channel.QueueDeclare(
		"supersandwich.sandwich_orders", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	json, err := o.ToJson()

	fmt.Println("######")
	fmt.Println("json:", string(json))
	fmt.Println("######")

	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	err = r.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        json,
		})

	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	log.Printf(" [x] Sent %s\n", string(json))

	return nil
}
