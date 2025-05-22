package internal

import (
	"context"
	"log"
	"vehicle-edge-telemetry/internal/domain"
	"vehicle-edge-telemetry/internal/gateways"
)

type VechicleEdgeTelemetry struct {
	context      context.Context
	logger       *log.Logger
	LogProcessor *gateways.LogProcessor
}

func NewVehicleEdgeTelemetry(ctx context.Context, logger *log.Logger, config domain.TelemetryServiceConfig) *VechicleEdgeTelemetry {
	v := &VechicleEdgeTelemetry{
		context: ctx,
		logger:  logger,
	}
	v.logger.Println("Initializing vehicle edge telemetry service")

	v.logger.Println("Initializing log processor")
	logProcessor, err := gateways.NewLogProcessor(ctx, logger, config.KafkaIngestionBrokers, config.KafkaIngestionTopic)
	if err != nil {
		v.logger.Fatalf("Error initializing log processor: %+v\n", err)
	}
	v.LogProcessor = logProcessor
	v.logger.Println("Log processor initialized successfully")

	return v
}

func (v *VechicleEdgeTelemetry) Start() {
	v.LogProcessor.Start()
}

func (v *VechicleEdgeTelemetry) Stop() {
	v.LogProcessor.Stop()
}
