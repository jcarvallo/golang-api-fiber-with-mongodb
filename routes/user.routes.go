package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jcarvallo/golang-api-fiber-with-mongodb/controllers"
)

func RoutesUser(App *fiber.App) {
	path := App.Group("/users")

	path.Get("", controllers.GetUsers)
	path.Post("", controllers.CreateUser)

	path.Get("/:id", controllers.GetUser)
	path.Put("/:id", controllers.UpdateUser)
	path.Delete("/:id", controllers.DeleteUser)
}
