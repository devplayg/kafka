package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"syscall"
	"github.com/devplayg/kafka/test1"
	"log"
)

func main() {
	consumer, err := sarama.NewConsumer(test1.Brokers, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	subscribe(test1.Topic, consumer)

	// Stop
	waitForSignals()
}

func subscribe(topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetNewest //get offset for the oldest message on the topic
	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				fmt.Printf("[offset-%d] %s\n", message.Offset, string(message.Value))
			}
		}(pc)
	}
}

//func messageReceived(message *sarama.ConsumerMessage) {
//	fmt.Println(string(message.Value))
//}

func waitForSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		log.Println("signal received, shutting down...")
	}
}
