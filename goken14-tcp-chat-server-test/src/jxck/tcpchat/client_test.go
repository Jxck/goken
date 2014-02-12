package tcpchat

import (
	"bytes"
	"testing"
)

func TestClient(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0))
	client := NewClient(buf)

	message := "test"
	client.Write(message)
	actual := bytes.Buffer(client.Conn)
	expected := message

	if actual != expected {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}
