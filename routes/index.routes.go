package routes

import "github.com/gofiber/fiber/v2"

func RoutesIndex(App *fiber.App) {
	RoutesUser(App)
}
