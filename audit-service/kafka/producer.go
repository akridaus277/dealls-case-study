package kafka

import (
	"audit-service/config"
	"audit-service/internal/messagebroker"
	"audit-service/utils"
	"context"

	"github.com/gin-gonic/gin"
)

func PublishAuditEvent(ctx *gin.Context, key []byte, value []byte) {
	// Setup Kafka publisher & controller
	brokers := []string{config.Env("KAFKA_BROKERS")}
	topic := config.Env("KAFKA_TOPIC")
	publisher := messagebroker.NewKafkaPublisher(brokers, topic)
	defer publisher.Close()

	err := publisher.PublishMessage(context.Background(), key, value)
	if err != nil {
		utils.SendResponse(ctx, utils.NewInternalServerErrorResponse("failed to publish to kafka"))

		return
	}
}
