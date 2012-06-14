package main

import (
	"../nsq"
	"log"
	"net"
)

var Protocols = map[int32]nsq.Protocol{}

func tcpClientHandler(clientConn net.Conn) error {
	client := nsq.NewServerClient(clientConn, clientConn.RemoteAddr().String())
	log.Printf("TCP: new client(%s)", client.String())
	go client.Handle(Protocols)
	return nil
}