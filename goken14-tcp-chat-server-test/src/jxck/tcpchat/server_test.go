package chat

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	port := ":3000"
	done := make(chan bool)
	go func() {
		server := NewServer()
		server.Serve(port)
		<-done
		server.Close()
	}()

	time.Sleep(time.Second)

	expected := "test\r\n"

	conn, _ := net.Dial("tcp", port)
	fmt.Fprintf(conn, expected)
	actual, _ := bufio.NewReader(conn).ReadString('\n')

	close(done)

	t.Logf("%q, %q", actual, expected)
	if actual != expected {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}

func TestMultiClient(t *testing.T) {
	port := ":3001"
	done := make(chan bool)
	go func() {
		server := NewServer()
		defer server.Close()
		server.Serve(port)
		<-done
	}()

	time.Sleep(time.Second)

	expected := "test\r\n"

	conn1, _ := net.Dial("tcp", port)
	conn2, _ := net.Dial("tcp", port)

	fmt.Fprintf(conn1, expected)
	actual1, _ := bufio.NewReader(conn1).ReadString('\n')
	actual2, _ := bufio.NewReader(conn2).ReadString('\n')

	close(done)

	t.Logf("%q, %q", actual1, actual2)
	if actual1 != actual2 {
		t.Errorf("\ngot  %v\nwant %v", actual1, actual2)
	}
}
