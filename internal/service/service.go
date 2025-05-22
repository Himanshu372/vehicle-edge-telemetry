package internal

import (
	"context"
	"log"
	"vehicle-edge-telemetry/internal/domain"

	"github.com/segmentio/kafka-go"
)

type VechicleEdgeTelemetry struct {
	context     context.Context
	logger      *log.Logger
	kafkaReader *kafka.Reader
}

func NewVehicleEdgeTelemetry(ctx context.Context, logger *log.Logger, config domain.TelemetryServiceConfig) *VechicleEdgeTelemetry {
	v := &VechicleEdgeTelemetry{
		context: ctx,
		logger:  logger,
	}
	v.logger.Println("Initializing vehicle edge telemetry service")

	v.logger.Println("Initializing Kafka reader")
	err := v.initKafka(config)
	if err != nil {
		v.logger.Fatalf("Error initializing Kafka reader: %+v\n", err)
	}
	v.logger.Println("Kafka reader initialized successfully")

	return v
}

func (v *VechicleEdgeTelemetry) initKafka(config domain.TelemetryServiceConfig) error {
	// Initialize Kafka reader
	v.kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  config.KafkaIngestionBrokers,
		Topic:    config.KafkaIngestionTopic,
		GroupID:  "vehicle-edge-group",
		MinBytes: config.KafkaIngestionMinBytes,
		MaxBytes: config.KafkaIngestionMaxBytes,
	})

	if err := v.kafkaReader.SetOffset(kafka.FirstOffset); err != nil {
		return err
	}

	return nil
}
