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

	topic = "test-offset"
)

func NewConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Producer.Return.Successes = true

	return config
}

func main() {
	// Check topic
	config := NewConfig()

	// Send Random
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = producer.Close() }()

	//partitions := getPartitions(config)
	//startConsumer(config, partitions)

	// Produce
	now := time.Now().Format(time.RFC3339)
	println(now)
	for i := 1; i <= 100; i++ {
		msg := &sarama.ProducerMessage{
			Topic:    topic,
			Value:    sarama.StringEncoder(now + " - " + strconv.Itoa(i)),
			Metadata: "test-meta",
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("partition: %d, offset: %d\n", partition, offset)
	}

}

//
//func startConsumer(config *sarama.Config, partitions []int32) {
//    master, err := sarama.NewConsumer(brokers, config)
//    if err != nil {
//        log.Fatal(err)
//    }
//    defer func() { _ = master.Close() }()
//    //master.ConsumePartition()
//
//    _, err = master.ConsumePartition(topic, 0, 1234)
//    if err != nil {
//        log.Fatal(err)
//    }
//}
//
//func runConsumer() {
//
//// Then: messages starting from offset 1234 are consumed.
//for i := 0; i < 10; i++ {
//    select {
//    case message := <-consumer.Messages():
//        assertMessageOffset(t, message, int64(i+1234))
//    case err := <-consumer.Errors():
//        t.Error(err)
//    }
//}
//
//safeClose(t, consumer)
//safeClose(t, master)
//broker0.Close()
//}

//func getPartitions(config *sarama.Config) []int32 {
//    admin, err := sarama.NewClusterAdmin(brokers, config)
//    if err != nil {
//        log.Fatal(err)
//    }
//    defer func() { _ = admin.Close() }()
//
//    metaList, err := admin.DescribeTopics([]string{topic})
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    if len(metaList) < 1 {
//        log.Fatal("no partitions")
//    }
//
//    var list []int32
//
//    for _, p := range metaList[0].Partitions {
//        list = append(list, p.ID)
//    }
//
//    return list
//}
