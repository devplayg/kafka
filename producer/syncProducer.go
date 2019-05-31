package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"time"
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
	//sendToLeader()
	//println("started")
	sendThruSarama()
}

func sendThruSarama() {

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
		return
	}

	i := 0
	for i < 1000 {
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Partition: -1,
			Value:     sarama.StringEncoder("test message: " + strconv.Itoa(i)),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("partition=%d, offset=%d\n", partition, offset)
		time.Sleep(1 * time.Second)
		i++

	}
}

//func sendToLeader() {
//    conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
//    if err != nil {
//        log.Fatal("failed to connect", err)
//        return
//    }
//    defer func() {
//        err := conn.Close()
//        if err != nil {
//            log.Fatal(err)
//        }
//    }()
//
//    _, err = conn.Write([]byte("message"))
//    if err != nil {
//        log.Fatal("failed to send message", err)
//        return
//    }
//}
