package controllers

import (
	"github.com/chaksack/payment-api/database"
	"github.com/chaksack/payment-api/models"
	"github.com/gofiber/fiber/v2"
)

func PaymentsAuthorization(c *fiber.Ctx) error {
	var payment []models.Payment
	database.Database.Db.Model(&models.Payment{})
}
