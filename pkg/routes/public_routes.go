package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaleemubarok/fiber/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/books", controllers.GetBooks)
	route.Get("/book/:id", controllers.GetBook)
	route.Get("/token/new", controllers.GetNewAccessToken)
	route.Post("/login", controllers.Login)

	route.Get("/accounts", controllers.GetAccounts)
	route.Get("/account/:id", controllers.GetAccount)
	route.Post("/account", controllers.CreateAccount)

}
