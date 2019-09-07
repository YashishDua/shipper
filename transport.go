package shipper

import (
	"fmt"
	"net"
)

type Transport struct {
	Type int
	Host string
	Port int
}

const (
	DefaultTransport int = 0
	TCPTransport     int = 1
)

func (transport *Transport) connect() (net.Conn, error) {
	conn, connErr := net.Dial("tcp", fmt.Sprintf("%s:%d", transport.Host, transport.Port))
	if connErr != nil {
		return nil, connErr
	}

	return conn, nil
}

func (transport *Transport) listen() (net.Listener, error) {
	listen, listenErr := net.Listen("tcp", fmt.Sprintf(":%d", transport.Port))
	if listenErr != nil {
		return nil, listenErr
	}

	//conn, _ := listen.Accept()

	return listen, nil
}
