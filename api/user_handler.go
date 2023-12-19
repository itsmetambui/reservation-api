package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsmetambui/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Tam",
		LastName: "Bui",
	}
	return c.JSON(user)
}


func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}