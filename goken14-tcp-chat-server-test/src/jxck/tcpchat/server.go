package tcpchat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Server struct {
	clients []*Client
}

func NewServer() *Server {
	server := &Server{
		clients: make([]*Client, 0),
	}
	return server
}

func (s *Server) Listen(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server starts at %v\n", port)

	accept := AcceptLoop(listener)
	broadcast := make(chan string)
	for {
		select {
		case client := <-accept:
			go client.ReadLoop(broadcast)
			s.clients = append(s.clients, client)
		case message := <-broadcast:
			go BroadCast(s.clients, message)
		default:
		}
	}
}

func AcceptLoop(listener net.Listener) chan *Client {
	accept := make(chan *Client)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			client := NewClient(conn)
			accept <- client
		}
	}()
	return accept
}

func BroadCast(clients []*Client, message string) {
	for _, client := range clients {
		go func(client *Client) {
			bw := bufio.NewWriter(client.Conn)
			bw.WriteString(message)
			bw.Flush()
		}(client)
	}
}
