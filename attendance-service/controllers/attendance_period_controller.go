package controllers

import (
	requestDto "attendance-service/dto/request"
	"attendance-service/services"
	"attendance-service/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddAttendancePeriod(c *gin.Context) {
	var input requestDto.AttendancePeriodRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}
	username, exists := c.Get("user")
	if !exists {
		utils.SendResponse(c, utils.NewUnauthorizedResponse("Unauthorized"))
		return
	}
	utils.SendResponse(c, services.AddAttendancePeriod(input.StartDate, input.EndDate, input.PayrollPeriodCode, fmt.Sprintf("%v", username)))
}

func GetAttendancePeriodPayroll(c *gin.Context) {
	var req requestDto.AttendancePeriodPayrollRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid Request"))
		return
	}

	utils.SendResponse(c, services.GetAttendancePeriodPayroll(req.PayrollPeriodCode))
}
