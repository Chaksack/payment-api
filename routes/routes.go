package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome to my payment api using gofiber framework")
	})

	//payments routes
	app.Post("/payments/authorization")

	//acocunts routes

}
