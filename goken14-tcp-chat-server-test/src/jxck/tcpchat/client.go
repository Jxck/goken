package tcpchat

import (
	"bufio"
	"fmt"
	"io"
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

func (c *Client) ReadLoop(broadcast chan string) {
	fmt.Printf("connect %+v\n", c)
	br := bufio.NewReader(c.Conn)
	for {
		message, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("dissconnect %+v\n", c.Conn)
			} else {
				log.Println(err)
			}
			return
		}
		fmt.Printf("%q\n", message)
		broadcast <- message
	}
}
