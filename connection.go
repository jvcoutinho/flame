package flame

import (
	"errors"

	"./handler"
	"./message"
)

// Connect connects to the following host and port.
func Connect(host string, port int, userID string) (*Channel, error) {

	handler := handler.New(host, port)
	if err := handler.Connect(); err != nil { // TODO: try again.
		return nil, errors.New("Connection could not be made, try again later")
	}

	channel := &Channel{handler}
	err := sendRegistrationRequest(userID, channel)
	return channel, err
}

func sendRegistrationRequest(identifier string, channel *Channel) error {
	request := message.NewRequest(identifier, message.Register, "", nil)
	if err := channel.call(request); err != nil {
		return err
	}
	return nil
}
