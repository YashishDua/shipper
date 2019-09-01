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

func (transport *Transport) connect() error {
	conn, connErr := net.Dial("tcp", fmt.Sprintf("%s:%d", transport.Host, transport.Port))
	if connErr != nil {
		return connErr
	}

	fmt.Println("TCP Connect: ", conn)
	return nil
}

func (transport *Transport) receive() error {
	listen, listenErr := net.Listen("tcp", fmt.Sprintf(":%d", transport.Port))
	if listenErr != nil {
		return listenErr
	}

	fmt.Println("TCP Receive: ", listen)
	return nil
}
