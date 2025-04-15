package topic

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"testing"
)

func TestDescribe(t *testing.T) {
	saramaConfig := sarama.NewConfig()
	brokers := []string{"127.0.0.1:9092"}
	client, err := sarama.NewClient(brokers, saramaConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	topic := "test1"
	// 获取 Topic 的所有分区
	partitions, err := client.Partitions(topic)
	if err != nil {
		log.Fatalf("Failed to get partitions: %v", err)
	}

	totalMessages := int64(0)

	// 遍历每个分区
	for _, partition := range partitions {
		// 获取最早偏移量（Oldest Offset）
		oldestOffset, err := client.GetOffset(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Printf("Failed to get oldest offset for partition %d: %v", partition, err)
			continue
		}

		// 获取最新偏移量（Latest Offset）
		latestOffset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Printf("Failed to get latest offset for partition %d: %v", partition, err)
			continue
		}

		// 计算分区的消息数量
		partitionMessages := latestOffset - oldestOffset
		totalMessages += partitionMessages

		fmt.Printf("Partition %d:\n", partition)
		fmt.Printf("  Oldest Offset: %d\n", oldestOffset)
		fmt.Printf("  Latest Offset: %d\n", latestOffset)
		fmt.Printf("  Messages: %d\n\n", partitionMessages)
	}

	fmt.Printf("Total messages in topic '%s': %d\n", topic, totalMessages)
}
