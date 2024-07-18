package kk

import (
	"encoding/json"
	"testing"
	"time"
)

//	func TestProducer(t *testing.T) {
//		// to produce messages
//		topic := "my-topic"
//		partition := 0
//
//		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
//		if err != nil {
//			log.Fatal("failed to dial leader:", err)
//		}
//
//		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
//		_, err = conn.WriteMessages(
//			kafka.Message{Value: []byte("one!")},
//			kafka.Message{Value: []byte("two!")},
//			kafka.Message{Value: []byte("three!")},
//		)
//		if err != nil {
//			log.Fatal("failed to write messages:", err)
//		}
//
//		if err := conn.Close(); err != nil {
//			log.Fatal("failed to close writer:", err)
//		}
//	}
//
//	func TestConsumer(t *testing.T) {
//		// to consume messages
//		topic := "my-topic"
//		partition := 0
//
//		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
//		if err != nil {
//			log.Fatal("failed to dial leader:", err)
//		}
//
//		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
//		batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
//
//		b := make([]byte, 10e3) // 10KB max per message
//		for {
//			n, err := batch.Read(b)
//			if err != nil {
//				break
//			}
//			fmt.Println(string(b[:n]))
//		}
//
//		if err := batch.Close(); err != nil {
//			log.Fatal("failed to close batch:", err)
//		}
//
//		if err := conn.Close(); err != nil {
//			log.Fatal("failed to close connection:", err)
//		}
//	}
func TestProducer(t *testing.T) {
	w := GetWriter("localhost:9092")
	m := make(map[string]string)
	m["projectCode"] = "1200"
	bytes, _ := json.Marshal(m)
	w.Send(LogData{
		Topic: "msproject_log",
		Data:  bytes,
	})
	time.Sleep(2 * time.Second)
}

func TestConsumer(t *testing.T) {
	GetReader([]string{"localhost:9092"}, "group1", "msproject_log")
	for {

	}
}
