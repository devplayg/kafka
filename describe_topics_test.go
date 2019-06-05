package kafka

import (
	"github.com/Shopify/sarama"
	"testing"
)

var (
	brokers = []string{
		"kafka1:9092",
		"kafka2:9092",
		"kafka3:9092",
	}

	config = NewConfig()

	topic = "test-topic"
)

func NewConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	return config
}

func TestCreateTopic(t *testing.T) {
	err := CreateTopic(brokers, config, topic)
	if err != nil {
		t.Error(err)
	}
}

func TestDescribeTopics(t *testing.T) {
	_, err := DescribeTopics(brokers, config, []string{topic})
	if err != nil {
		t.Error(err)
	}
}
