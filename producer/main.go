package main

import (
	"fmt"
	"log"
	"net/http"
	"rabbitmqtest/producer/services/rabbitmq"

	"github.com/gin-gonic/gin"
)

func main() {
	rabbitmq.NewRabbitMqStore()

	defer rabbitmq.RabbitMQClient.CloseConnection()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/sandwich", makeSandwichOrderPOST)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
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
