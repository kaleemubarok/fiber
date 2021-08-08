package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaleemubarok/fiber/pkg/configs"
	"github.com/kaleemubarok/fiber/pkg/middleware"
	"github.com/kaleemubarok/fiber/pkg/routes"
	"github.com/kaleemubarok/fiber/pkg/utils"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServer(app)
}
