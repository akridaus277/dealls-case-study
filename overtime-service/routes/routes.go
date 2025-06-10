package routes

import (
	"overtime-service/controllers"
	"overtime-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.POST("/overtime", controllers.SubmitOvertime)
	}

	apiMicroservice := r.Group("/services")
	{
		apiMicroservice.POST("/get-overtime-summary", controllers.GetOvertimeSummary)
	}

}
