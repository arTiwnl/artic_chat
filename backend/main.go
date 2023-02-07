package main

import (
	"fmt"
	"net/http"
 
	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)
// definindo o endpoint da lib websocket

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
	    fmt.Fprintf(w, "%+v\n", err)
	}
 
	client := &websocket.Client{
	    Conn: conn,
	    Pool: pool,
	}
 
	pool.Register <- client
	client.Read()
 }

func setupRoutes() {
	// mapeando o endpoint /ws para a funçao a seguir
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
	})
}



// rota de saida //
func main() {
	fmt.Println("Distributed Chat App v0.06")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
 }