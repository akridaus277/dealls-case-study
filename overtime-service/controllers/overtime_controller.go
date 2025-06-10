package controllers

import (
	"fmt"
	requestDto "overtime-service/dto/request"
	"overtime-service/services"
	"overtime-service/utils"

	"github.com/gin-gonic/gin"
)

func SubmitOvertime(c *gin.Context) {
	username, exists := c.Get("user")
	if !exists {
		utils.SendResponse(c, utils.NewUnauthorizedResponse("Unauthorized"))
		return
	}

	var input requestDto.OvertimeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}

	utils.SendResponse(c, services.SubmitOvertime(fmt.Sprintf("%v", username), input.Hours))
}

func GetOvertimeSummary(c *gin.Context) {
	var req requestDto.OvertimeSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse(fmt.Sprintf("Invalid Request : %s", err)))
		return
	}

	utils.SendResponse(c, services.GetOvertimeSummary(req.PeriodStart, req.PeriodEnd))
}
