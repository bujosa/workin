package models

import (
	_"fmt"
)

type MessageModel struct {
	
}

var messages []string = []string{"Action Complete", "Action Imcomplete"}

func (messagemodel MessageModel)Success() (string, error) {
	
    return messages[0], nil
}

func (messagemodel MessageModel)Denied() (string, error) {
	
    return messages[1], nil
}


