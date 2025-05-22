package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"vehicle-edge-telemetry/internal/domain"
	internal "vehicle-edge-telemetry/internal/service"
)

func main() {
	fmt.Println("Service started")
	// temp config
	conf := domain.TelemetryServiceConfig{
		KafkaIngestionBrokers:  []string{"localhost:9092"},
		KafkaIngestionTopic:    "vehicle-edge-topic",
		KafkaIngestionMinBytes: 10e3,
	}
	ctx := context.Background()
	logger := log.New(os.Stdout, "vehicle-edge-telemetry: ", log.LstdFlags)
	// Initialize the service
	v := internal.NewVehicleEdgeTelemetry(ctx, logger, conf)
	fmt.Println("Service stopped")

}
