package websocket

import (
	"fmt"

)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// eu tirei a pontuçao de em branco de todos os for a seguir: ` client, _ := `
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Conexoes ativas: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "Novo usuario conectado !"})
			}
			break
		case client := <- pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Conexoes ativas: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "Se desconectou"})
			}
			break
		case message := <- pool.Broadcast:
			fmt.Println("Mensagem enviada a todos da sessão")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}



		}
	}
}