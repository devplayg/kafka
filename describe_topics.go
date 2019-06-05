package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func CreateTopic(brokers []string, config *sarama.Config, topic string) error {
	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		return err
	}
	defer admin.Close()

	err = admin.CreateTopic(topic, &sarama.TopicDetail{NumPartitions: 5, ReplicationFactor: 3}, false)
	if err != nil {
		return err
	}

	err = admin.Close()
	if err != nil {
		return err
	}

	return nil
}

func DescribeTopics(brokers []string, config *sarama.Config, topics []string) ([]*sarama.TopicMetadata, error) {
	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		return nil, err
	}
	defer admin.Close()

	metaList, err := admin.DescribeTopics(topics)
	if err != nil {
		return nil, err
	}

	for _, meta := range metaList {
		fmt.Printf("Topic: %s, PartitionCount:%d\n", meta.Name, len(meta.Partitions))
		for _, p := range meta.Partitions {
			fmt.Printf("\tPartition: %d, Leader: %d, Replicas: %v, Isr: %v\n", p.ID, p.Leader, p.Replicas, p.Isr)
		}
	}

	//Topic:my-topic  PartitionCount:5        ReplicationFactor:3     Configs:
	//Topic: my-topic Partition: 0    Leader: 2       Replicas: 2,1,3 Isr: 3,1,2
	//Topic: my-topic Partition: 1    Leader: 2       Replicas: 3,2,1 Isr: 2,1
	//Topic: my-topic Partition: 2    Leader: 2       Replicas: 1,3,2 Isr: 2,1
	//Topic: my-topic Partition: 3    Leader: 2       Replicas: 2,3,1 Isr: 3,1,2
	//Topic: my-topic Partition: 4    Leader: 2       Replicas: 3,1,2 Isr: 2,1

	//Topic:my-topic  PartitionCount:5        ReplicationFactor:3     Configs:
	//Topic: my-topic Partition: 0    Leader: 2       Replicas: 2,1,3 Isr: 3,1,2
	//Topic: my-topic Partition: 1    Leader: 2       Replicas: 3,2,1 Isr: 2,1,3
	//Topic: my-topic Partition: 2    Leader: 2       Replicas: 1,3,2 Isr: 2,1,3
	//Topic: my-topic Partition: 3    Leader: 2       Replicas: 2,3,1 Isr: 3,1,2
	//Topic: my-topic Partition: 4    Leader: 2       Replicas: 3,1,2 Isr: 2,1,3
	return metaList, nil
}
