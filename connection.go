package flame

import (
	"errors"

	"./handler"
)

// Connect connects to the following host and port.
func Connect(host string, port int) (*Channel, error) {
	handler := handler.New(host, port)
	if err := handler.Connect(); err != nil { // TODO: try again.
		return nil, errors.New("Connection could not be made, try again later")
	}
	return &Channel{handler}, nil
}
