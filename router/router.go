package router

import (
	"apirest/controller"
	"apirest/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	// Book
	feed := api.Group("/feed")
	feed.Get("/", controller.GetBooks)
	feed.Get("/:id", controller.GetBook)

	feed.Use(middleware.JWTProtected)
	feed.Post("/", controller.CreateBook)
	feed.Patch("/:id", controller.UpdateBook)
	feed.Delete("/:id", controller.DeleteBook)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)
}
