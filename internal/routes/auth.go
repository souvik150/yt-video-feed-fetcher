package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/souvik150/Fampay-Backend-Assignment/internal/controllers"
)

func AuthRoutes(group fiber.Router) {
	userGroup := group.Group("/user")
	userGroup.Post("/login", controllers.LoginUser)
	userGroup.Post("/verify", controllers.VerifyOTP)
	userGroup.Post("/resend", controllers.ResendOTP)
	userGroup.Post("/signup", controllers.SignupUser)
	userGroup.Post("/refresh", controllers.RefreshToken)
}
