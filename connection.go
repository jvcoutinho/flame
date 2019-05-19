package flame

import (
	"./handler"
)

// Connect connects to the following host and port.
func Connect(host string, port int) (*Channel, error) {
	handler := handler.New(host, port)
	if err := handler.Connect(); err != nil {
		return nil, err
	}
	return &Channel{handler}, nil
}
