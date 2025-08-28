package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rabbitmqtest/producer/services/rabbitmq"

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
	fmt.Println("da port:", port)

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
			"message": "pong",
		})
	})

	router.POST("/sandwich", makeSandwichOrderPOST)

	router.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080
}

func makeSandwichOrderPOST(c *gin.Context) {
	var r rabbitmq.SandwichRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(r)

	err := rabbitmq.RabbitMQClient.MakeASandwichRequest(r)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send coffee order to RabbitMQ queue"})
		return
	}

	c.JSON(http.StatusOK, r)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
