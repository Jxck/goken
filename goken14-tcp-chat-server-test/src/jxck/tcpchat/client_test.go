package chat

import (
	"bytes"
	"testing"
)

type rwcMock struct {
	bytes.Buffer
	closed bool
}

func (m *rwcMock) Close() error {
	m.closed = true
	return nil
}

func TestReadLoop(t *testing.T) {
	id := 1
	conn := &rwcMock{}
	broadCastChan := make(chan string)
	client := NewClient(id, conn, broadCastChan, nil)

	expected := "message\r\n"
	client.Conn.Write([]byte(expected))

	actual := <-broadCastChan

	if actual != expected {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}

func TestLeaveChan(t *testing.T) {
	id := 1
	conn := &rwcMock{}
	broadCastChan := make(chan string)
	leaveChan := make(chan *Client)
	client := NewClient(id, conn, broadCastChan, leaveChan)

	client.Conn.Close()

	actual := <-client.LeaveChan
	expected := client
	t.Logf("\n%#v\n%#v", expected, client)

	if actual != expected {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}
