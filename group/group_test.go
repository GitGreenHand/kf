package group

import (
	"fmt"
	"github.com/IBM/sarama"
	"testing"
)

func TestGroup(t *testing.T) {
	saramaConfig := sarama.NewConfig()
	brokers := []string{"127.0.0.1:9092"}
	admin, err := sarama.NewClusterAdmin(brokers, saramaConfig)
	if err != nil {
		t.Fatal(err)
	}
	groups, err := admin.ListConsumerGroups()

	for groupId, _ := range groups {
		fmt.Println(groupId)
	}
	if err != nil {
		t.Fatal(err)
	}

}
