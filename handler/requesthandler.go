package handler

import (
	"net"
	"strconv"
)

type ClientRequestHandler struct {
	host string
	port int
	conn net.Conn
}

func New(host string, port int) *ClientRequestHandler {
	return &ClientRequestHandler{host, port, nil}
}

func (crh *ClientRequestHandler) Connect() error {
	conn, err := net.Dial("tcp", crh.host+":"+strconv.Itoa(crh.port))
	if err != nil {
		return err
	}
	crh.conn = conn
	return nil
}

func (crh *ClientRequestHandler) Send(data []byte) {
	crh.conn.Write(data)
}
