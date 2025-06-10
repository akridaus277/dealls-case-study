package controllers

import (
	"fmt"
	requestDto "reimbursement-service/dto/request"
	"reimbursement-service/services"
	"reimbursement-service/utils"

	"github.com/gin-gonic/gin"
)

func SubmitReimbursement(c *gin.Context) {
	username, exists := c.Get("user")
	if !exists {
		utils.SendResponse(c, utils.NewUnauthorizedResponse("Unauthorized"))
		return
	}

	var input requestDto.ReimbursementRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}

	utils.SendResponse(c, services.SubmitReimbursement(fmt.Sprintf("%v", username), input.Amount, input.Description))
}

func GetReimbursementSummary(c *gin.Context) {
	var req requestDto.ReimbursementSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse(fmt.Sprintf("Invalid Request : %s", err)))
		return
	}

	utils.SendResponse(c, services.GetReimbursementSummary(req.PeriodStart, req.PeriodEnd))
}
