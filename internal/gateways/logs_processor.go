package gateways

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const (
	// MaxBytes is the maximum number of bytes to read from Kafka
	MaxBytes = 10e6 // 10MB
)

type LogProcessor struct {
	ctx         context.Context
	logger      *log.Logger
	kafkaReader *kafka.Reader
	brokers     []string
	topic       string
}

func NewLogProcessor(ctx context.Context, logger *log.Logger, brokers []string, topic string) (*LogProcessor, error) {
	if len(brokers) == 0 || topic == "" {
		return nil, fmt.Errorf("brokers and topic must be provided: %v, %s", brokers, topic)
	}
	lp := &LogProcessor{
		ctx:     ctx,
		logger:  logger,
		brokers: brokers,
		topic:   topic,
	}

	lp.logger.Println("Initializing Kafka reader")
	reader, err := lp.initKafka()
	if err != nil {
		lp.logger.Fatalf("Error initializing Kafka reader: %+v\n", err)
	}
	lp.kafkaReader = reader
	lp.logger.Println("Kafka reader initialized successfully")
	return lp, nil
}

func (lp *LogProcessor) initKafka() (*kafka.Reader, error) {
	// Initialize Kafka reader
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  lp.brokers,
		Topic:    lp.topic,
		GroupID:  "vehicle-edge-group",
		MaxBytes: MaxBytes,
	})

	if err := kafkaReader.SetOffset(kafka.FirstOffset); err != nil {
		return nil, err
	}

	return kafkaReader, nil
}

func (lp *LogProcessor) processLog(msg kafka.Message) {
	// Process the log message
	lp.logger.Printf("Processing log message: %s\n", string(msg.Value))
	// Here you can add your logic to process the log message
}

func (lp *LogProcessor) Start() {
	lp.logger.Println("Starting log processor")
	for {
		select {
		case <-lp.ctx.Done():
			lp.logger.Println("Stopping log processor")
			return
		default:
			msg, err := lp.kafkaReader.ReadMessage(lp.ctx)
			if err != nil {
				lp.logger.Printf("Error reading message: %v\n", err)
				continue
			}
			lp.processLog(msg)
		}
	}
}

func (lp *LogProcessor) Stop() {
	lp.logger.Println("Stopping log processor")
	// Close the Kafka reader
	if err := lp.kafkaReader.Close(); err != nil {
		lp.logger.Printf("Error closing Kafka reader: %v\n", err)
	}
	lp.logger.Println("Log processor stopped")
}
