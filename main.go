package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jcarvallo/golang-api-fiber-with-mongodb/routes"
)

var App *fiber.App

func main() {
	App := fiber.New()

	routes.RoutesIndex(App)

	App.Listen(":5000")
}
