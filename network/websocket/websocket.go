package websocket

import (
	"net/http"

	"github.com/kr/pretty"
	"github.com/gorilla/websocket"

	"github.com/D0K-ich/KanopyService/logs"
)

var log = logs.NewLog()

var ws_client = websocket.Upgrader{
	ReadBufferSize	: 1024,				//1 Kb
	WriteBufferSize	: 1024,				//1 Kb
	CheckOrigin		: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn, err_chan chan error) {
	var message_type int
	var data []byte
	var err error

	for {
		if message_type, data, err = conn.ReadMessage(); err != nil {pretty.Println(err);return}

		pretty.Println(string(data))
		if err = conn.WriteMessage(message_type, []byte("Lol")); err != nil {pretty.Println(err);return}
	}
}


func NewWsConnection() (err error) {
	log.Info("Create new ws connection")

	setupRoutes()
	if err = http.ListenAndServe(":8080", nil); err != nil {return}

	return
}

