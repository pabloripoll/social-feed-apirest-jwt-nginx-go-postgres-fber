package router

import (
	memberAuth "apirest/domain/member/controller/auth"
	//memberProfile "apirest/domain/member/controller/profile"
	feedPost "apirest/domain/feed/controller"
	"apirest/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Get("/", func(res *fiber.Ctx) error {
		return res.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Social Feed V1 - Index.",
		})
	})

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", memberAuth.Login)
	auth.Post("/register", memberAuth.Register)

	// Member Account
	//account := api.Group("/account")
	//account.get("/", memberProfile.Profile)

	// Feed
	feed := api.Group("/feed")
	feed.Get("/", feedPost.GetFeeds)
	feed.Get("/:id", feedPost.GetFeed)

	feed.Use(middleware.JWTProtected)
	feed.Post("/", feedPost.CreateFeed)
	feed.Patch("/:id", feedPost.UpdateFeed)
	feed.Delete("/:id", feedPost.DeleteFeed)
}
