package tcpchat

import (
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Client struct {
	Conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	client := &Client{
		Conn: conn,
	}
	return client
}
