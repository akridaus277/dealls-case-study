package routes

import (
	"api-gateway-service/middlewares"
	"api-gateway-service/proxy"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Group microservice auth
	authServiceGroup := r.Group("/services/auth")
	authServiceGroup.Any("/*proxyPath", proxy.ProxyAuthService)
	// Group api auth
	authGroup := r.Group("/api/auth")
	authGroup.Any("/*proxyPath", proxy.ProxyAuthService)

	// Group microservice employee
	employeeServiceGroup := r.Group("/services/employees")
	employeeServiceGroup.Any("/*proxyPath", proxy.ProxyEmployeeService)
	// Group api employee , apply middleware JWT untuk semua /api/employees/*
	employeeGroup := r.Group("/api/employees", middlewares.JWTAuthMiddleware())
	employeeGroup.Any("/*proxyPath", proxy.ProxyEmployeeService)

	// Group microservice attendance
	attendanceServiceGroup := r.Group("/services/attendance")
	attendanceServiceGroup.Any("/*proxyPath", proxy.ProxyAttendanceService)
	// Group api attendance , apply middleware JWT untuk semua /api/attendance/*
	attendanceGroup := r.Group("/api/attendance", middlewares.JWTAuthMiddleware())
	attendanceGroup.Any("/*proxyPath", proxy.ProxyAttendanceService)

	// Group microservice overtime
	overtimeServiceGroup := r.Group("/services/overtime")
	overtimeServiceGroup.Any("/*proxyPath", proxy.ProxyOvertimeService)
	// Group api overtime , apply middleware JWT untuk semua /api/overtime/*
	overtimeGroup := r.Group("/api/overtime", middlewares.JWTAuthMiddleware())
	overtimeGroup.Any("/*proxyPath", proxy.ProxyOvertimeService)

	// Group microservice reimbursement
	reimbursementServiceGroup := r.Group("/services/reimbursement")
	reimbursementServiceGroup.Any("/*proxyPath", proxy.ProxyReimbursementService)
	// Group api reimbursement , apply middleware JWT untuk semua /api/reimbursement/*
	reimbursementGroup := r.Group("/api/reimbursement", middlewares.JWTAuthMiddleware())
	reimbursementGroup.Any("/*proxyPath", proxy.ProxyReimbursementService)

	// Group microservice payroll
	payrollServiceGroup := r.Group("/services/payroll")
	payrollServiceGroup.Any("/*proxyPath", proxy.ProxyPayrollService)
	// Group api payroll , apply middleware JWT untuk semua /api/payroll/*
	payrollGroup := r.Group("/api/payroll", middlewares.JWTAuthMiddleware())
	payrollGroup.Any("/*proxyPath", proxy.ProxyPayrollService)

	auditGroup := r.Group("/audit")
	auditGroup.Any("/*proxyPath", proxy.ProxyAuditService)
}
