package controllers

import (
	"audit-service/internal/messagebroker"
	"audit-service/kafka"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuditController struct {
	Publisher *messagebroker.KafkaPublisher
}

type AuditEvent struct {
	Event     string `json:"event"`
	Service   string `json:"service"`
	RequestID string `json:"request_id"`
	IP        string `json:"ip"`
	Timestamp string `json:"timestamp"`
}

func PublishAuditEvent(c *gin.Context) {
	// Ambil userID dari header (dummy sementara)
	userIDStr := c.Request.Header.Get("X-User-ID")
	userID, _ := strconv.Atoi(userIDStr)

	var payload AuditEvent
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	// Inject userID ke log kalau mau (opsional)
	payload.Service += " [user_id:" + strconv.Itoa(userID) + "]"

	msgBytes, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error encoding JSON"})
		return
	}

	kafka.PublishAuditEvent(c, nil, msgBytes)

	c.JSON(http.StatusOK, gin.H{"message": "event published to kafka"})
}
