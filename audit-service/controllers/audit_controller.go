package controllers

import (
	"audit-service/kafka"
	"audit-service/models"
	"audit-service/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuditEvent struct {
	Action   string `json:"action"`
	Metadata string `json:"metadata"`
	Service  string `json:"service"`
}

func PublishAuditEvent(c *gin.Context) {
	requestID := c.GetHeader("X-Request-ID")
	ipAddress := c.ClientIP()

	username, exists := c.Get("user")
	if !exists {
		utils.SendResponse(c, utils.NewUnauthorizedResponse("Unauthorized"))
		return
	}

	var payload AuditEvent
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	// Buat log lengkap
	logEntry := models.AuditLog{
		RequestID: requestID,
		IPAddress: ipAddress,
		Action:    payload.Action,
		Metadata:  payload.Metadata,
		Service:   payload.Service,
		Username:  fmt.Sprintf("%v", username),
	}

	msgBytes, err := json.Marshal(logEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error encoding JSON"})
		return
	}

	kafka.PublishAuditEvent(c, nil, msgBytes)
	c.JSON(http.StatusOK, gin.H{"message": "event published to kafka"})
}
