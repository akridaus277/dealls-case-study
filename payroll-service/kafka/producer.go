package kafka

import (
	"context"
	"payroll-service/config"
	"payroll-service/internal/messagebroker"

	"github.com/gin-gonic/gin"
)

func PublishAuditEvent(ctx *gin.Context, key []byte, value []byte) error {
	// Setup Kafka publisher & controller
	brokers := []string{config.Env("KAFKA_BROKERS")}
	topic := config.Env("KAFKA_TOPIC")
	publisher := messagebroker.NewKafkaPublisher(brokers, topic)
	defer publisher.Close()

	err := publisher.PublishMessage(context.Background(), key, value)
	if err != nil {
		return err
	}
	return nil
}
