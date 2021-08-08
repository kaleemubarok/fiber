package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	readTimeOutSecondCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeOutSecondCount),
	}
}
