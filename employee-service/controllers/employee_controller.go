package controllers

import (
	requestDto "employee-service/dto/request"
	"employee-service/models"
	"employee-service/services"
	"employee-service/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterEmployee(c *gin.Context) {
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}

	utils.SendResponse(c, services.RegisterEmployee(emp))
	return
}

func ListEmployees(c *gin.Context) {
	utils.SendResponse(c, services.ListEmployees())
	return
}

func GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}

	utils.SendResponse(c, services.GetEmployeeByID(id))
	return
}

func UpdateEmployee(c *gin.Context) {
	var data map[string]interface{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request id"))
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request body"))
		return
	}

	utils.SendResponse(c, services.UpdateEmployee(id, data))
	return
}

func DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request id"))
		return
	}

	utils.SendResponse(c, services.DeleteEmployee(id))
	return

}

func GetEmployeeByNip(c *gin.Context) {
	var data requestDto.GetByNip

	if err := c.ShouldBindJSON(&data); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request body"))
		return
	}

	utils.SendResponse(c, services.GetEmployeeByNIP(data.Nip))
	return
}
