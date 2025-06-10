package routes

import (
	"employee-service/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/employee")
	{
		api.GET("/:id", controllers.GetEmployee)
		api.PUT("/:id", controllers.UpdateEmployee)
		api.DELETE("/:id", controllers.DeleteEmployee)
	}

	apiMicroservice := r.Group("/services/employee")
	{
		apiMicroservice.POST("/get-by-nip", controllers.GetEmployeeByNip)
		apiMicroservice.GET("/get-all-employee", controllers.ListEmployees)
	}
}
