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
			"message": "pong DING DONG)))",
		})
	})

	router.POST("/sandwich", makeSandwichOrderPOST)
	router.GET("/menu", getMenuGET)

	router.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080
}


// @Summary      Creates a sandwich order
// @Description  Creates a sandwich order
// @Tags         create-sandwich
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Router       /users/{id} [get]
func makeSandwichOrderPOST(c *gin.Context) {
	var requestData rabbitmq.SandwichRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(requestData)

	err := rabbitmq.RabbitMQClient.MakeASandwichRequest(requestData)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send sandwich order to RabbitMQ queue"})
		return
	}

	c.JSON(http.StatusOK, requestData)
}

type Menu struct {
	FranchiseName string
	Menu map[string]float32
	BaseCurrency string
}

func getMenuGET(c *gin.Context) {
	n := Menu {
		FranchiseName: "SuperSandwich",
		Menu: map[string]float32{
			"Vegan EXTREME": 5.3,
			"Meat EXTREME": 15.0,
			"Burrito EXTREME": 12.0,
		},
		BaseCurrency: "euro",
	}

 	c.JSON(http.StatusOK, n)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
