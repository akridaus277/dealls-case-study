package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func proxyToTarget(c *gin.Context, targetHost string) {
	target, err := url.Parse(targetHost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid target host"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Overwrite Director untuk menjaga Authorization header dan path
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		// Set ulang path sesuai parameter proxyPath
		req.URL.Path = c.Param("proxyPath")

		// Copy header dari request asli ke proxy
		req.Header = c.Request.Header.Clone()
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func ProxyAuthService(c *gin.Context) {
	proxyToTarget(c, os.Getenv("AUTH_SERVICE_URL"))
}

func ProxyEmployeeService(c *gin.Context) {
	proxyToTarget(c, os.Getenv("EMPLOYEE_SERVICE_URL"))
}

func ProxyAttendanceService(c *gin.Context) {
	proxyToTarget(c, os.Getenv("ATTENDANCE_SERVICE_URL"))
}

func ProxyAuditService(c *gin.Context) {
	proxyToTarget(c, os.Getenv("AUDIT_SERVICE_URL"))
}

func ProxyOvertimeService(c *gin.Context) {
	proxyToTarget(c, os.Getenv("OVERTIME_SERVICE_URL"))
}

func ProxyReimbursementService(c *gin.Context) {
	proxyToTarget(c, os.Getenv("REIMBURSEMENT_SERVICE_URL"))
}

func ProxyPayrollService(c *gin.Context) {
	proxyToTarget(c, os.Getenv("PAYROLL_SERVICE_URL"))
}
