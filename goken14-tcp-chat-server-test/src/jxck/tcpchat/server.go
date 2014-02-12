package tcpchat

/**
 * 小さい方の実装。
 * Accept と Read と Broadcast をそれぞれ別の goroutine ループで回している。
 * Connection はスライスで持っている。
 **/

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

type Server struct {
}

func NewServer() *Server {
	return new(Server)
}

func (s *Server) Listen(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server starts at %v\n", port)

	accept := AcceptLoop(listener)
	clients := make([]*Client, 0)
	broadcast := make(chan string)
	for {
		select {
		case conn := <-accept:
			client := NewClient(conn)
			go ReadLoop(client, broadcast)
			clients = append(clients, client)
		case message := <-broadcast:
			go BroadCast(clients, message)
		default:
		}
	}
}

func AcceptLoop(listener net.Listener) chan net.Conn {
	accept := make(chan net.Conn)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			accept <- conn
		}
	}()
	return accept
}

func ReadLoop(client *Client, broadcast chan string) {
	fmt.Printf("connect %v\n", client)
	br := bufio.NewReader(client.Conn)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("dissconnect %v\n", client.Conn)
			} else {
				log.Println(err)
			}
			return
		}
		message := string(line) + "\n"
		log.Printf("%q\n", message)
		broadcast <- string(message)
	}
}

func BroadCast(clients []*Client, message string) {
	for _, client := range clients {
		go func(conn net.Conn) {
			bw := bufio.NewWriter(conn)
			bw.WriteString(message)
			bw.Flush()
		}(client.Conn)
	}
}
