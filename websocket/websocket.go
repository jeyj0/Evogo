package websocket

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	gorillaWs "github.com/gorilla/websocket"
)

type Socket struct {
	clients []*gorillaWs.Conn
	message string
}

func NewInit() {
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

	go func(socket *Socket) {
		counter := 0
		for {
			socket.message = strconv.Itoa(counter)
			time.Sleep(2000 * time.Millisecond)
			counter++
		}
	}(&socket)

	panic(http.ListenAndServe(":2345", nil))
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

func (socket Socket) broadcast() {
	for _, client := range socket.clients {
		go func(client *gorillaWs.Conn) {
			err := client.WriteJSON(socket.message)
			if err != nil {
				socket.removeClient(client)
			}
		}(client)
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
