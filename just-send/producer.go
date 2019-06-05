package main

import (
    "github.com/Shopify/sarama"
    "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
    "log"
)

var (
    brokers = []string{
        "kafka1:9092",
        "kafka2:9092",
        "kafka3:9092",
    }
)

func main() {
    config := sarama.NewConfig()
    config.Version = sarama.V2_2_0_0

    admin, err := sarama.NewClusterAdmin(brokers, config)
    if err != nil {
        log.Fatal(err)
        return
    }
    defer admin.Close()

    topics := []string{"my-topic"}
    topicMetaList, err := admin.DescribeTopics(topics)
    if err != nil {
        log.Fatal(err)
        return
    }

    for _, meta := range topicMetaList {
        fmt.Printf("%smeta.Name)
        //spew.Dump(meta)
    }

}

func

func createToipc() {
    //kafka.TopicConfig{}
    //retention := "-1"
    //
    //
    //req := &sarama.CreateTopicsRequest{
    //   TopicDetails: map[string]*sarama.TopicDetail{
    //       "topic": {
    //           NumPartitions:     5,
    //           ReplicationFactor: 3,
    //           ConfigEntries: map[string]*string{
    //               "retention.ms": &retention,
    //           },
    //       },
    //   },
    //   Timeout: 100 * time.Millisecond,
    //}
    //



    //sarama.TopicDetail{
    //    NumPartitions:     -1,
    //    ReplicationFactor: -1,
    //    ReplicaAssignment: map[int32][]int32{
    //        0: []int32{0, 1, 2},
    //    },
    //    ConfigEntries: map[string]*string{
    //        "retention.ms": &retention,
    //    },
    //}
}

//
//func NewConsumer2(brokers, topics, group_id string, handler CommonHandler) {
//    // Init config, specify appropriate version
//    config := sarama.NewConfig()
//    config.Version = sarama.V0_10_2_0
//    config.Consumer.Return.Errors = true
//    config.Consumer.Offsets.Initial = sarama.OffsetOldest
//
//    // Start with a client
//    brokersList := strings.Split(brokers, ",")
//    client, err := sarama.NewClient(brokersList, config)
//    if err != nil {
//        panic(err)
//    }
//    // gorputine 退出时，关闭客户端
//    defer func() { _ = client.Close() }()
//
//    // Start a new consumer group
//    group, err := sarama.NewConsumerGroupFromClient(group_id, client)
//    if err != nil {
//        panic(err)
//    }
//    defer func() { _ = group.Close() }()
//
//    // Track errors
//    go func() {
//        for err := range group.Errors() {
//            fmt.Println("KAFKA |group Errors=", err)
//        }
//    }()
//
//    // Iterate over consumer sessions.
//    ctx := context.Background()
//    topicsList := strings.Split(topics, ",")
//
//    for {
//        err := group.Consume(ctx, topicsList,
//            exampleConsumerGroupHandler{
//                handler: handler,
//            })
//        if err != nil {
//            fmt.Println("KAFKA |Consume err=", err)
//        }
//    }
//
//    return
//}