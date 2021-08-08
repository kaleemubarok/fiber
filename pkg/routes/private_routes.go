package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaleemubarok/fiber/app/controllers"
	"github.com/kaleemubarok/fiber/pkg/middleware"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook)
}
