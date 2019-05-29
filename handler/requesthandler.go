package handler

import (
	"net"
	"strconv"
)

type RequestHandler struct {
	host string
	port int
	conn net.Conn
}

func New(host string, port int) *RequestHandler {
	return &RequestHandler{host, port, nil}
}

func (crh *RequestHandler) Connect() error {
	conn, err := net.Dial("tcp", crh.host+":"+strconv.Itoa(crh.port))
	if err != nil {
		return err
	}
	crh.conn = conn
	return nil
}

func (crh *RequestHandler) Send(data []byte) error {
	_, err := crh.conn.Write(data)
	return err
}

func (handler *RequestHandler) Receive() ([]byte, error) {
	byteMsg := make([]byte, 2048)
	read, err := handler.conn.Read(byteMsg)
	if err != nil {
		return nil, err
	}
	return byteMsg[:read], nil
}
