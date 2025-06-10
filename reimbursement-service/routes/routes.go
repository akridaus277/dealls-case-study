package routes

import (
	"reimbursement-service/controllers"
	"reimbursement-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.POST("/reimbursement", controllers.SubmitReimbursement)
	}

	apiMicroservice := r.Group("/services")
	{
		apiMicroservice.POST("/get-reimbursement-summary", controllers.GetReimbursementSummary)
	}
}
