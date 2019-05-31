package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

var (
	brokers = []string{
		"kafka1:9092",
		"kafka2:9092",
		"kafka3:9092",
	}

	topic = "my-topic"
)

func main() {
	//kafka.DialPartition()

	//conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
	conn, err := kafka.DialPartition(context.Background(), "tcp", brokers[0], topic, 0)
	if err != nil {
		log.Fatal("failed to connect", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = conn.Write([]byte("message"))
	if err != nil {
		log.Fatal("failed to send message", err)
		return
	}

	//kafka.NewConn()
	//kafka.DialPartition(context.Background() "tcp", "127.0.0.1:8001", topic, 0)
	//
	//
	//conn, err := kafka.Dial("tcp", "127.0.0.1:8001")
	//if err != nil {
	//    log.Fatal("failed to connect", err)
	//}
	//defer conn.Close()
	//
	//_, err = conn.Write([]byte("hello from go"))
	//if err != nil {
	//    log.Fatal("failed to send message", err)
	//}

	//kafka.Dial()
	//conn := kafka.NewClient(brokers[0], brokers[1], brokers[2])

	//writerConfig := kafka.WriterConfig{
	//    Brokers: brokers,
	//    Topic: topic,
	//}
	//
	//writer := kafka.NewWriter(writerConfig)
	//writer.WriteMessages()

	//kafka.Conn{
	//
	//}
	//w := kafka.NewWriter(kafka.WriterConfig{
	//   Brokers: brokers,
	//   Topic:   topic,
	//   Balancer: &kafka.LeastBytes{},
	//})
	//defer w.Close()
	//
	//err := w.WriteMessages(context.Background(),
	//   kafka.Message{
	//       Key:   []byte("Key-A"),
	//       Value: []byte("Hello World!"),
	//   },
	//)
	//
	//if err != nil {
	//   log.Fatal(err)
	//}
}
