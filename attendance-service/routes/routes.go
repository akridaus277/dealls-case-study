package routes

import (
	"attendance-service/controllers"
	"attendance-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	admin := r.Group("/admin", middlewares.JWTAuthAdminMiddleware())
	{
		admin.POST("/attendance-period", controllers.AddAttendancePeriod)
	}

	api := r.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.POST("/attendance", controllers.SubmitAttendance)
	}

	apiMicroservice := r.Group("/services")
	{
		apiMicroservice.POST("/get-attendance-summary", controllers.GetAttendanceSummary)
		apiMicroservice.POST("/get-attendance-period", controllers.GetAttendancePeriodPayroll)
	}

}
