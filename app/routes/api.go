package routes

import (
	"v1/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes initializes authentication routes
func SetupRoutes(app *fiber.App) {
    auth := app.Group("/user")

    // Get All User
    auth.Get("/all", controllers.UserController)


    // Registration route
    auth.Post("/register", controllers.RegisterController)

	

    // Login route
    // auth.Post("/login", controllers.LoginController)

    // Logout route
    // auth.Post("/logout", controllers.LogoutController)
}