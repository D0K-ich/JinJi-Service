package websocket

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gorilla/websocket"
)

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	var web_socket *websocket.Conn
	var err error

	if web_socket, err = ws_client.Upgrade(w, r, nil); err != nil {log.Warn("Err on ws", zap.Any("err", err));return}

	if err = web_socket.WriteMessage(1, []byte("Hi Client!")); err != nil {log.Warn("Err on ws", zap.Any("err", err));return}

	reader(web_socket, nil)
	return
}