package main

import (
	"context"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("Start")
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.Open(filepath.Join(currDir + "/../data/sim_data/can_log_simulated.csv"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:8001"},
		Topic:    "simulator-logs",
		Balancer: &kafka.LeastBytes{},
		Dialer:   dialer,
	})

	recordChan := make(chan []string)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		for record := range recordChan {
			val := struct {
				DLC         string
				DataPayload string
				Timestamp   string
			}{record[0], record[2], record[3]}
			valBytes, err := json.Marshal(&val)
			if err != nil {
				fmt.Printf("error when marshaling val: %+v, skipping record\n", val)
				continue
			}
			kafkaWriter.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte(record[1]),
					Value: []byte(valBytes),
				},
			)
		}
	}()

	for _, record := range records {
		recordChan <- record
	}
	wg.Wait()
	fmt.Println("End")
}
