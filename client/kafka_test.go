package client

import (
	"fmt"
	"github.com/IBM/sarama"
	"testing"
)

func TestKafkaClient(t *testing.T) {

	saramaConfig := sarama.NewConfig()
	brokers := []string{"127.0.0.1:9092"}
	admin, err := sarama.NewClusterAdmin(brokers, saramaConfig)
	if err != nil {
		t.Fatal(err)
	}
	topics, err := admin.ListTopics()
	if err != nil {
		t.Fatal(err)
	}
	for topic := range topics {
		fmt.Println(topic)
	}
}
