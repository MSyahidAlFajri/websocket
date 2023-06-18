package typestruct

import "github.com/gofiber/websocket/v2"

type Messagee struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Clientt struct {
	Username string
	Conn     *websocket.Conn
}

type ChatRoomm struct {
	Clients    []*Clientt
	Register   chan *Clientt
	Unregister chan *Clientt
	Broadcast  chan Messagee
}

type NewMessagee struct {
	Id      string
	Message string
}
