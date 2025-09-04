package internal

import (
	"fmt"
	"net/http"
	"rabbitmqtest/shared/rabbitmq"

	"github.com/gin-gonic/gin"
)

// @Summary      Creates a sandwich order
// @Description  Creates a sandwich order
// @Tags         create-sandwich
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Router       /users/{id} [get]
func MakeSandwichOrderPOST(c *gin.Context) {
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

func GetMenuGET(c *gin.Context) {
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