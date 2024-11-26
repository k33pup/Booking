//mock delivery system

package deliverysystem

import (
	"encoding/json"
	"fmt"

	logger "github.com/k33pup/Booking.git/pkg/logger"
)

type Message struct {
	Str string `json:"message"`
}

type IDeliverySystem struct {
	SendMessage func([]byte) error
}

type MockDeliverySystem struct {
}

func (*MockDeliverySystem) SendMessage(bytes []byte) error {
	msg := Message{}
	err := json.Unmarshal(bytes, &msg)
	if err != nil {
		logger.WriteLog(err.Error()) //TODO
	}
	fmt.Println(msg.Str)
	return nil
}

func NewMockDeliverySystem() *MockDeliverySystem {
	d_system := MockDeliverySystem{}
	return &d_system
}
