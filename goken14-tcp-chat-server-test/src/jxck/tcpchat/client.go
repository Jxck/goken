package tcpchat

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Client struct {
}

func NewClient() *Client {
	return new(Client)
}
