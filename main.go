package main

import (
	"github.com/chaksack/payment-api/database"
	"github.com/chaksack/payment-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()
	app := fiber.New()

	//new addition

	routes.Setup(app)

	app.Listen(":3000")
}
