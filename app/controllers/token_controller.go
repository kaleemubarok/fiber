package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaleemubarok/fiber/pkg/utils"
)

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status
func GetNewAccessToken(c *fiber.Ctx) error {
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
