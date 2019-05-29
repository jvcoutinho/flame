package flame

import (
	"errors"

	"./handler"
	"./message"
)

// Connect connects to the following host and port.
func Connect(host string, port int, userID string) (*Channel, error) {

	requestHandler := handler.New(host, port)
	itemHandler := handler.New(host, port)
	if tryConnection(requestHandler) != nil || tryConnection(itemHandler) != nil {
		return nil, errors.New("Connection could not be made, try again later")
	}

	channel := &Channel{userID, requestHandler, itemHandler, make(map[string]*chan interface{})}
	go channel.receiveItem()
	err := sendRegistrationRequest(userID, channel)
	return channel, err
}

func tryConnection(handler *handler.RequestHandler) error {
	return handler.Connect() // TODO: try again.
}

func sendRegistrationRequest(identifier string, channel *Channel) error {
	request := message.NewRequest(identifier, message.Register, "", nil, 0)
	if err := channel.call(request); err != nil {
		return err
	}
	return nil
}
