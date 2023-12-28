package routes

import (
	"v1/app/controllers"
	auth "v1/app/controllers/auth"

	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes initializes authentication routes
func SetupRoutes(app *fiber.App) {
    user := app.Group("/user")
    // Get All User
    user.Get("/list", controllers.GetUserList)
    // user.Get("/:id", controllers.GetUserById)


    // Registration route
    user.Post("/register", auth.Register)

	

    // Login route
    // auth.Post("/login", controllers.LoginController)

    // Logout route
    // auth.Post("/logout", controllers.LogoutController)
}