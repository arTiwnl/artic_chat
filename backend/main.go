package main

import (
	"fmt"
	"net/http"

	"https://github.com/arTiwnl/artic_chat/tree/master/backend/pkg/websocket"
)

// definindo o endpoint da lib websocket

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit")
	conn, err := websocket.Upgrade(w,r)
	if err != nil {
		fmt.Fprintf(w, "%+\n", err)
	}
	client := &websocket.Client {
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
	fmt.Println("Chat App v.04")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}