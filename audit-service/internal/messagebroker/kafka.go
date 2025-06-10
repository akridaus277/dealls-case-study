package messagebroker

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaSubscriber struct {
	reader *kafka.Reader
}

type KafkaPublisher struct {
	writer *kafka.Writer
}

// ----------- SUBSCRIBER (CONSUMER) ------------
func NewKafkaSubscriber(brokers []string, topic, groupID string) *KafkaSubscriber {
	return &KafkaSubscriber{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     brokers,
			Topic:       topic,
			GroupID:     groupID,
			StartOffset: kafka.FirstOffset,
		}),
	}
}

func (s *KafkaSubscriber) ReadMessage(ctx context.Context) ([]byte, error) {
	msg, err := s.reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}
	return msg.Value, nil
}

func (s *KafkaSubscriber) Close() error {
	return s.reader.Close()
}

// ----------- PUBLISHER (PRODUCER) ------------
func NewKafkaPublisher(brokers []string, topic string) *KafkaPublisher {
	return &KafkaPublisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *KafkaPublisher) PublishMessage(ctx context.Context, key, value []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

func (p *KafkaPublisher) Close() error {
	return p.writer.Close()
}
