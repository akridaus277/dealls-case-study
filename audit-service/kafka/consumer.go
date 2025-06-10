package kafka

import (
	"audit-service/internal/messagebroker"
	"audit-service/models"
	service "audit-service/services"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type KafkaConfig struct {
	Brokers []string
	Topic   string
	GroupID string
}

func StartConsumer(svc *service.AuditService, ctx context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found. Using system environment variables.")
	}

	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	topic := os.Getenv("KAFKA_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP_ID")

	subscriber := messagebroker.NewKafkaSubscriber(brokers, topic, groupID)
	defer subscriber.Close()

	log.Println("[Kafka] Listening to topic:", topic)

	for {
		data, err := subscriber.ReadMessage(ctx)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		var logData models.AuditLog
		if err := json.Unmarshal(data, &logData); err != nil {
			log.Println("Invalid audit log format:", err)
			continue
		}

		log.Printf("[AUDIT] %s | %s | %s | %s\n", logData.Service, logData.Action, logData.IPAddress, logData.RequestID)

		if err := svc.Log(&logData); err != nil {
			log.Println("Failed to save audit log:", err)
		}
	}
}
