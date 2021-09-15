package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func consumer() {
	conn, err := kafka.DialLeader(context.Background(),"tcp", "192.168.33.13:9092", "simple-topic", 0)
	if err != nil{
		log.Fatal("failed to dial leader :", err)
	}
	conn.SetReadDeadline(time.Now().Add(10 *time.Second))
	batch := conn.ReadBatch(10e3, 1e6)
	b := make([]byte, 10e3)
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}
	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal(" failed to close connection:", err)
	}
}

func Reader() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"192.168.33.13:9092"},
		Topic:     "simple-topic",
		Partition: 0,
		MinBytes:  10, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(2)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func main() {
	//consumer()
	Reader()
}
