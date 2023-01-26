package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

// definindo o endpoint da lib websocket

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	// mapeando o endpoint /ws para a funçao a seguir
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v.01")
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}