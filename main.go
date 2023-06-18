package main

import (
	"log"

	"github.com/MSyahidAlFajri/websocket/module"
	"github.com/MSyahidAlFajri/websocket/url"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.Run()

	app := fiber.New()
	url.Web(app)
	log.Fatal(app.Listen(musik.Dangdut()))
}
