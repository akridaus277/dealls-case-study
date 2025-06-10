package routes

import (
	"authentication-service/controllers"
	"authentication-service/internal/feignclient"
	"authentication-service/services"

	_ "authentication-service/cmd/docs"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Inisialisasi client dan service
	employeeClient := feignclient.NewEmployeeClient()
	authService := services.NewAuthService(employeeClient)
	authController := controllers.NewAuthController(authService)

	api := r.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)

	}

}
