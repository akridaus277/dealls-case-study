package routes

import (
	"payroll-service/controllers"
	"payroll-service/internal/feignclient"
	"payroll-service/middlewares"
	"payroll-service/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	employeeClient := feignclient.NewEmployeeClient()
	attendanceClient := feignclient.NewAttendanceClient()
	overtimeClient := feignclient.NewOvertimeClient()
	reimbursementClient := feignclient.NewReimbursementClient()
	payrollService := services.NewPayrollService(employeeClient, attendanceClient, overtimeClient, reimbursementClient)
	payrollController := controllers.NewPayrollController(payrollService)

	admin := r.Group("/admin", middlewares.JWTAuthAdminMiddleware())
	{
		admin.POST("/payroll/run", payrollController.RunPayroll)
		admin.POST("/payslip/report-excel", payrollController.GetPayslipAllEmployee)
	}
	api := r.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.POST("/payslip", payrollController.GetPayslip)
	}
}
