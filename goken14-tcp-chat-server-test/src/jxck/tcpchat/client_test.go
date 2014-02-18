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
	id := 0
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
