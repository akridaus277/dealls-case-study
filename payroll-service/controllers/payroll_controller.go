package controllers

import (
	"fmt"
	"payroll-service/services"
	"payroll-service/utils"

	"github.com/gin-gonic/gin"
)

type PayrollController struct {
	PayrollService *services.PayrollService
}

// Constructor
func NewPayrollController(payrollService *services.PayrollService) *PayrollController {
	return &PayrollController{
		PayrollService: payrollService,
	}
}

func (pc *PayrollController) RunPayroll(c *gin.Context) {
	var input struct {
		PayrollCode string `json:"PayrollCode"`
	}

	_, exists := c.Get("user")
	if !exists {
		utils.SendResponse(c, utils.NewUnauthorizedResponse("Unauthorized"))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}

	utils.SendResponse(c, pc.PayrollService.RunPayroll(input.PayrollCode))

}

func (pc *PayrollController) GetPayslip(c *gin.Context) {
	var input struct {
		PayrollCode string `json:"PayrollCode"`
	}

	nip, exists := c.Get("user")
	if !exists {
		utils.SendResponse(c, utils.NewUnauthorizedResponse("Unauthorized"))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}

	pc.PayrollService.GeneratePayslipPDF(c, input.PayrollCode, fmt.Sprintf("%v", nip))
}

func (pc *PayrollController) GetPayslipAllEmployee(c *gin.Context) {
	var input struct {
		PayrollCode string `json:"PayrollCode"`
	}

	_, exists := c.Get("user")
	if !exists {
		utils.SendResponse(c, utils.NewUnauthorizedResponse("Unauthorized"))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}

	pc.PayrollService.ExportPayslipExcel(c, input.PayrollCode)
}
