package main

import (
	"fmt"
	"os"
	"rabbitmqtest/services/producer/internal"
	"rabbitmqtest/shared/rabbitmq"

	"github.com/gin-gonic/gin"
)

var version = "0.1.8-development"

func main() {
	fmt.Println("Rabbit producer version", version)
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
