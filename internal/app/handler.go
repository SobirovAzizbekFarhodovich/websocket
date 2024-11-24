package app

import (
	"fmt"
	"net/http"
	"websocket/config"
	"websocket/internal/domain"
	"websocket/internal/service"
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := config.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket ulanishi xatosi:", err)
		return
	}
	service.AddClient(conn)

	for {
		var msg domain.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Xabarni o'qishda xato:", err)
			service.RemoveClient(conn)
			break
		}

		if msg.IsJoin {
			fmt.Printf("Foydalanuvchi qo'shildi: %s\n", msg.Username)
		} else if msg.FileName != "" && msg.FileData != "" {
			fmt.Printf("Foydalanuvchi %s fayl yubordi: %s\n", msg.Username, msg.FileName)
		} else {
			fmt.Printf("Foydalanuvchi %s yubordi: %s\n", msg.Username, msg.Content)
		}

		service.BroadcastMessage(msg)
	}
}
