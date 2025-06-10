package routes

import (
	"audit-service/controllers"
	"audit-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.POST("/test/publish", controllers.PublishAuditEvent)
	}

}
