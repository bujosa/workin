package entities

import (
	"fmt"
)

type Message struct {
	Result string
}

type Messages []Message

func (message Message) ToString() string {
	return fmt.Sprintf("Result: %s",
		message.Result)
}