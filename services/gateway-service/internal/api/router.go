package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/api/v1"
	"github.com/hexolan/panels/gateway-service/internal/api/handlers"
)

func RegisterRoutes(app *fiber.App) {
	apiV1 := app.Group("/v1")

	// Panel Service Routes
	panelV1 := apiV1.Group("/panels")
	panelV1.Post("/", v1.CreatePanel) // todo: check permissions
	panelV1.Get("/:name", v1.GetPanelByName)
	panelV1.Patch("/:name", handlers.AuthMiddleware, v1.UpdatePanelByName) // todo: check permissions
	panelV1.Delete("/:name", handlers.AuthMiddleware, v1.DeletePanelByName) // todo: check permissions

	// Post Service Routes
	postV1 := apiV1.Group("/posts")
	postV1.Patch("/:id", handlers.AuthMiddleware, v1.UpdatePost) // todo: check permissions
	postV1.Delete("/:id", handlers.AuthMiddleware, v1.DeletePost) // todo: check permissions
	
	panelV1.Post("/:panel_name", v1.CreatePanelPost)
	panelV1.Get("/:panel_name/posts", v1.GetPanelPosts)
	panelV1.Get("/:panel_name/posts/:id", v1.GetPanelPost)

	// User Service Routes
	userV1 := apiV1.Group("/users")
	userV1.Post("/", v1.UserSignup)
	userV1.Get("/me", handlers.AuthMiddleware, v1.GetCurrentUser)
	userV1.Get("/:username", handlers.AuthMiddleware, v1.GetUserByUsername)
	userV1.Delete("/delete", handlers.AuthMiddleware, v1.DeleteCurrentUser)

	// Auth Service Routes
	// todo: changing passwords
	authV1 := apiV1.Group("/auth")
	authV1.Post("/login", v1.LoginWithPassword)

	// Comment Service Routes
	commentV1 := postV1.Group("/:post_id/comments")
	commentV1.Get("/", v1.GetPostComments)
	commentV1.Post("/", handlers.AuthMiddleware, v1.CreateComment)
	commentV1.Patch("/:id", handlers.AuthMiddleware, v1.UpdateComment) // todo: check permissions
	commentV1.Delete("/:id", handlers.AuthMiddleware, v1.DeleteComment) // todo: check permissions
}