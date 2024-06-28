package network

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kr/pretty"
	"net/http"
	"testing"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {pretty.Println(err);return}

		// print out that message for clarity
		pretty.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {pretty.Println(err)}
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {pretty.Println(err)}
	pretty.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {pretty.Println(err)}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func TestWs(t *testing.T) {
	fmt.Println("Hello World")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
