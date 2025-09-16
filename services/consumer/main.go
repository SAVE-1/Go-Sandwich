package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

var version = "0.1.2-development"

/* 

	http://localhost:15672/#/ 

*/
func main() {
	fmt.Println("Rabbit consumer version", version)
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}

	urlPostGresql := os.Getenv("POSTGRESQL_URL")
	if urlPostGresql == "" {
		// postgresql-orders
		// urlPostGresql = "user=postgres password=example dbname=postgres sslmode=disable"
		urlPostGresql = "postgres://postgres:example@localhost/postgres?sslmode=disable"
	}

	fmt.Println(urlPostGresql)

	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"supersandwich.sandwich_orders", // name
		true,                            // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
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
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}


	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			uglyInsert(urlPostGresql)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func uglyInsert(conn string) {
	fmt.Println("uglyInsert()")

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	name := "jee"

	rows, err := db.Query("SELECT name FROM test WHERE name = $1", name)

	for rows.Next() {
		var num string

		err = rows.Scan(&num)
		if err != nil {
			panic(err)
		}

		fmt.Println(num)
	}

	sqlStatement := `
INSERT INTO test (name)
VALUES ($1)
RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, "test@test.random").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
