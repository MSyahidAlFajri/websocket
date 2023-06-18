package url

import (
	"github.com/MSyahidAlFajri/websocket/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.GetIP)
	page.Get("/ws", websocket.New(controller.Websocket))

}
