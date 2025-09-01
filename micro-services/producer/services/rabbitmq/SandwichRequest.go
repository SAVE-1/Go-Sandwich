package rabbitmq

import (
	"encoding/json"
	"fmt"
)

type SandwichRequest struct {
	Count int    `json:"Count" binding:"required"`
	Name  string `json:"Name" binding:"required"`
	Type  int    `json:"Type"`
}

func (s SandwichRequest) GetCount() int {
	return s.Count
}

func (s SandwichRequest) GetName() string {
	return s.Name
}

func (s SandwichRequest) GetType() int {
	return s.Type
}

func (s SandwichRequest) ToJson() ([]byte, error) {
	b, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s SandwichRequest) String() string {
	return fmt.Sprintf("SandwichRequest{Count:%d, Name:%q, Type:%d}", s.Count, s.Name, s.Type)
}

type ObjectRequest interface {
	GetCount() int
	GetName() string
	GetType() int
	ToJson() ([]byte, error)
}
