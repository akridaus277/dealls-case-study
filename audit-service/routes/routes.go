package routes

import (
	"audit-service/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/test/publish", controllers.PublishAuditEvent)
}
