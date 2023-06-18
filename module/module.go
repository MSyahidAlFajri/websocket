package module

import (
	"fmt"
	"log"

	"github.com/MSyahidAlFajri/websocket/typestruct"
)

func NewChatRoom() *typestruct.ChatRoomm {
	return &typestruct.ChatRoomm{
		Clients:    make([]*typestruct.Clientt, 0),
		Register:   make(chan *typestruct.Clientt),
		Unregister: make(chan *typestruct.Clientt),
		Broadcast:  make(chan typestruct.Messagee),
	}
}
func BroadcastMessage(message typestruct.Messagee) {
	NewChatRoom().Broadcast <- message
}

func Run() {
	for {
		select {
		case client := <-NewChatRoom().Register:
			NewChatRoom().Clients = append(NewChatRoom().Clients, client)
			go BroadcastMessage(typestruct.Messagee{
				Username: "Server",
				Content:  fmt.Sprintf("User %s joined the chat", client.Username),
			})
		case client := <-NewChatRoom().Unregister:
			for i, c := range NewChatRoom().Clients {
				if c == client {
					NewChatRoom().Clients = append(NewChatRoom().Clients[:i], NewChatRoom().Clients[i+1:]...)
					go BroadcastMessage(typestruct.Messagee{
						Username: "Server",
						Content:  fmt.Sprintf("User %s left the chat", client.Username),
					})
					break
				}
			}
		case message := <-NewChatRoom().Broadcast:
			for _, client := range NewChatRoom().Clients {
				go func(c *typestruct.Clientt) {
					if err := c.Conn.WriteJSON(message); err != nil {
						log.Println("Error broadcasting message:", err)
					}
				}(client)
			}
		}
	}
}
