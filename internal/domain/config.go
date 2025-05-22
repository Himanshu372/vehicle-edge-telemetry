package domain

type TelemetryServiceConfig struct {
	KafkaIngestionBrokers  []string `config:"env=KAFKA_NODES;usage=comma-separated list of brokers"`
	KafkaIngestionTopic    string   `config:"env=KAFKA_TOPIC;usage=topic to read from"`
	KafkaIngestionMinBytes int      `config:"env=KAFKA_MIN_BYTES;default= 10e3;usage=minimum bytes to read"`
	KafkaIngestionMaxBytes int      `config:"env=KAFKA_MAX_BYTES;default= 10e6;usage=maximum bytes to read"`
}
