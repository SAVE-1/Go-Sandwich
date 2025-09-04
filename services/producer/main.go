package main

import (
	"fmt"
	"log"
	"os"
	"rabbitmqtest/services/producer/internal"
	"rabbitmqtest/shared/rabbitmq"

	"github.com/gin-gonic/gin"
)

func main() {
	// set APP_ENV=production
	// $env:PRODUCER_HOST_PORT = "8081"
	// setx PRODUCER_HOST_PORT 8081
	// setx RABBITMQ_URL amqp://guest:guest@localhost:5672/
	// docker run -d --hostname my-rabbit --name some-rabbit-2 rabbitmq:3-management
	// docker run -d --hostname my-rabbit -p 15672:15672 --name some-rabbit rabbitmq:3-management
	// docker run -d --hostname my-rabbit --name some-rabbit -p 8080:15672 5672:5672 rabbitmq:3-management
	// docker run -d --hostname my-rabbit --name some-rabbit -p 8080:15672 -p 5672:5672 rabbitmq:3-management
	port := os.Getenv("PRODUCER_HOST_PORT")
	fmt.Println(".env port:", port)

	if port == "" {
		port = "8080"
	}

	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}
	fmt.Println("url:", url)

	if err := rabbitmq.NewRabbitMqStore(url); err != nil {
		fmt.Println(err)
		return
	}

	defer rabbitmq.RabbitMQClient.CloseConnection()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong DING DONG)))",
		})
	})

	router.POST("/sandwich", internal.MakeSandwichOrderPOST)
	router.GET("/menu", internal.GetMenuGET)

	router.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
