package main

import (
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {

	//topicConfig := kafka.TopicConfig{
	//    Topic: "event",
	//    NumPartitions: 5,
	//    ReplicationFactor:3,
	//}
	//topicConfig.Topic
	//
	//kafka.DialLeader()
	//
	//conn, err := kafka.Dial("localhost:8001", )
	//
	//conn, err := Dial("tcp", "localhost:9092")
	//if err != nil {
	//    t.Error("bad conn")
	//    return
	//}
	//defer conn.Close()
	//
	//_, err = conn.createTopics(createTopicsRequestV0{
	//    Topics: []createTopicsRequestV0Topic{
	//        {
	//            Topic:             topic,
	//            NumPartitions:     int32(partitions),
	//            ReplicationFactor: 1,
	//        },
	//    },
	//    Timeout: int32(30 * time.Second / time.Millisecond),
	//})

	topicConfig := kafka.TopicConfig{
		Topic:             "event",
		NumPartitions:     5,
		ReplicationFactor: 3,
	}

	conn, err := kafka.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Fatal("err1", err)
	}

	err = conn.CreateTopics(topicConfig)
	if err != nil {
		log.Fatal("err2", err)
	}

	println("topic created: ", topicConfig.Topic)

}
