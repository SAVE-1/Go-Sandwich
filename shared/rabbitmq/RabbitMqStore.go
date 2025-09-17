package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

func NewRabbitMqStore(connectionString string) error {
	connection, err := amqp.Dial(connectionString)

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

// func (r *RabbitMQStore) MakeASandwichRequest(o []ObjectRequest) error {
func (r *RabbitMQStore) MakeASandwichRequest(sandwich SandwichRequest) error {
	q, err := r.Channel.QueueDeclare(
		"supersandwich.sandwich_orders", // name
		true,                            // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bodyBytes, err := json.Marshal(sandwich)

	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = r.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
		})

	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	log.Printf(" [x] Sent %s\n", string(bodyBytes))

	return nil
}
