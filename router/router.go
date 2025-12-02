package router

import (
	"apirest/controller"
	"apirest/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	// Feed
	feed := api.Group("/feed")
	feed.Get("/", controller.GetFeeds)
	feed.Get("/:id", controller.GetFeed)

	feed.Use(middleware.JWTProtected)
	feed.Post("/", controller.CreateFeed)
	feed.Patch("/:id", controller.UpdateFeed)
	feed.Delete("/:id", controller.DeleteFeed)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)
}
