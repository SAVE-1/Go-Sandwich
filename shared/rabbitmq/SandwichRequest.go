package rabbitmq

import (
	"encoding/json"
	"fmt"
)

type SingularSandwich struct {
	Count int    `json:"Count" binding:"required"`
	Name  string `json:"Name" binding:"required"`
	Type  int    `json:"Type"`
}

// SandWich order creation POST struct
type SandwichRequest struct {
	Sandwiches []SingularSandwich
}

func (s SingularSandwich) GetCount() int {
	return s.Count
}

func (s SingularSandwich) GetName() string {
	return s.Name
}

func (s SingularSandwich) GetType() int {
	return s.Type
}

func (s SingularSandwich) ToJson() ([]byte, error) {
	b, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s SingularSandwich) String() string {
	return fmt.Sprintf("SingularSandwich{Count:%d, Name:%q, Type:%d}", s.Count, s.Name, s.Type)
}
