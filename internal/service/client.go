package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"websocket/internal/domain"
)

var clients []*websocket.Conn

func AddClient(conn *websocket.Conn) {
	clients = append(clients, conn)
}

func RemoveClient(conn *websocket.Conn) {
	for i, client := range clients {
		if client == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

func BroadcastMessage(msg domain.Message) {
	for _, client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			fmt.Println("Xabarni yuborishda xato:", err)
			RemoveClient(client)
		}
	}
}
