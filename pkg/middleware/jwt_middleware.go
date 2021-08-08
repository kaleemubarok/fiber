package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtMiddleWare "github.com/gofiber/jwt/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleWare.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtMiddleWare.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
