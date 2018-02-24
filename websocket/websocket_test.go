package websocket

import (
	"testing"

	gorillaWs "github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestRemoveClient1(t *testing.T) {
	// given
	client1 := &gorillaWs.Conn{}
	client2 := &gorillaWs.Conn{}
	socket := Socket{clients: []*gorillaWs.Conn{client1, client2}}

	// when
	socket.removeClient(client1)

	// then
	assert.Equal(t, []*gorillaWs.Conn{client2}, socket.clients)
}

func TestRemoveClient2(t *testing.T) {
	// given
	client1 := &gorillaWs.Conn{}
	client2 := &gorillaWs.Conn{}
	socket := Socket{clients: []*gorillaWs.Conn{client1, client2}}

	// when
	socket.removeClient(client2)

	// then
	assert.Equal(t, []*gorillaWs.Conn{client1}, socket.clients)
}
