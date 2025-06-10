package controllers

import (
	requestDto "attendance-service/dto/request"
	"attendance-service/services"
	"attendance-service/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SubmitAttendance(c *gin.Context) {
	username, exists := c.Get("user")
	if !exists {
		utils.SendResponse(c, utils.NewUnauthorizedResponse("Unauthorized"))
		return
	}

	utils.SendResponse(c, services.SubmitAttendance(fmt.Sprintf("%v", username)))
}

func GetAttendanceSummary(c *gin.Context) {
	var req requestDto.AttendanceSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse(fmt.Sprintf("Invalid Request : %s", err)))
		return
	}

	utils.SendResponse(c, services.GetAttendanceSummary(req.PeriodStart, req.PeriodEnd))
}
