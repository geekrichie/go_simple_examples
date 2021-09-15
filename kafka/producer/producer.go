package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func Producer() {
	conn , err := kafka.DialLeader(context.Background(), "tcp", "192.168.33.13:9092", "simple-topic", 0)
	if err != nil {
		log.Fatal("failed to dial leader: ", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer", err)
	}
}

func listTopic() {
	conn, err := kafka.Dial("tcp", "192.168.33.13:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
		fmt.Println(p.Leader)
	}
	for k := range m {
		fmt.Println(k)
	}
}

func main() {
	//Producer()
	listTopic()
}