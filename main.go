package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/itsmetambui/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", api.HandleGetUsers)
	apiv1.Get("/users/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine!"})
}