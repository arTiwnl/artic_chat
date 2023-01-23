package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// verificando se há conexao
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a rota que ira ouvir as mensagens 

func reader(conn *websocket.Conn) {
	for {
	// lendo as mensagens
	messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return 
		} 
		fmt.Println(string(p))

			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}

	}
}

// definindo o endpoint da lib websocket

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// atualizando o websocket e a conexao

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Servidor simples")
	})
	// mapeando o endpoint /ws para a funçao a seguir
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v.01")
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}