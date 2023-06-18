package controller

import (
	"log"

	"github.com/MSyahidAlFajri/websocket/module"
	"github.com/MSyahidAlFajri/websocket/typestruct"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Websocket(c *websocket.Conn) {
	username := c.Query("username")
	client := &typestruct.Clientt{
		Username: username,
		Conn:     c,
	}
	module.NewChatRoom().Register <- client

	defer func() {
		module.NewChatRoom().Unregister <- client
		c.Close()
	}()

	for {
		var message typestruct.Messagee
		err := c.ReadJSON(&message)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		module.BroadcastMessage(message)
	}
}

func GetIP(c *fiber.Ctx) error {
	getip := musik.GetIPaddress()
	return c.JSON(getip)
}
