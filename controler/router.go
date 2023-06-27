package controler

import (
	"gayandn/handler"
	"gayandn/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Use(middleware.Protected())
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetUser)
	user.Post("/CreateUser", handler.CreateUser)
	user.Patch("/UpdateUser/:id", handler.UpdateUser)
	user.Delete("/DeleteUser/:id", handler.DeleteUser)

}
