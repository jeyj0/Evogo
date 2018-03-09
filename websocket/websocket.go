package websocket

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	gorillaWs "github.com/gorilla/websocket"
)

type Socket struct {
	clients []*gorillaWs.Conn
	Message string
}

func Init() *Socket {
	socket := Socket{clients: []*gorillaWs.Conn{}}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(&socket, w, r)
	})

	go func(socket *Socket) {
		for {
			socket.broadcast()
			time.Sleep(200 * time.Millisecond)
		}
	}(&socket)

	go func() {
		panic(http.ListenAndServe(":2345", nil))
	}()

	return &socket
}

var upgrader = gorillaWs.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWebSocket(socket *Socket, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	socket.clients = append(socket.clients, conn)
}

func (socket *Socket) broadcast() {
	for _, client := range socket.clients {
		go func(socket *Socket, client *gorillaWs.Conn) {
			err := client.WriteJSON(socket.Message)
			if err != nil {
				socket.removeClient(client)
			}
		}(socket, client)
	}
}

func (socket *Socket) removeClient(client *gorillaWs.Conn) {
	for i := len(socket.clients) - 1; i >= 0; i-- {
		if socket.clients[i] == client {
			socket.clients = append(socket.clients[:i], socket.clients[i+1:]...)
			fmt.Println("Removed client")
			return
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}
