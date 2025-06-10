package kafka

import (
	"authentication-service/config"
	"authentication-service/internal/messagebroker"
	"context"

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
