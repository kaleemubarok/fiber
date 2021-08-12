package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaleemubarok/fiber/app/queries"
	"github.com/kaleemubarok/fiber/platform/database"
)

func OpenDBConnection() (*queries.AccountQueries, error) {
	db, err := database.PostgresSQLConnetion()
	if err != nil {
		return nil, err
	}

	return &queries.AccountQueries{db}, nil
}

func GetAccounts(c *fiber.Ctx) error {
	db, err := OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	accounts, err := db.GetAccounts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "account were not found" + err.Error(),
			"count":   0,
			"account": nil,
		})
	}
	for _, accont := range accounts {
		accont.Salt = ""
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"count":    len(accounts),
		"accounts": accounts,
	})
}
