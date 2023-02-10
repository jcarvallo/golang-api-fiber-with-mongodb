package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jcarvallo/golang-api-fiber-with-mongodb/models"
	"github.com/jcarvallo/golang-api-fiber-with-mongodb/services"
	"github.com/jcarvallo/golang-api-fiber-with-mongodb/utils"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetUsers()
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(users)
}
func GetUser(c *fiber.Ctx) error {
	user, err := services.GetUser(c.Params("id"))
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(user)
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)
	res, err := services.CreateUser(user)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(res)
}
func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)
	res, err := services.UpdateUser(user, c.Params("id"))
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(res)
}

func DeleteUser(c *fiber.Ctx) error {
	err := services.DeleteUser(c.Params("id"))
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleNotContent(c, "Delete user")
}
