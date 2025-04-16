package topic

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"kf/client"
	"log"
	"time"
)

func NewMessageCmd() *cobra.Command {

	createCmd := &cobra.Command{
		Use:     "message",
		Short:   "message topic to kafka",
		Aliases: []string{"m"},
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.GetClient()
			if err != nil {
				log.Fatalf("Error get client: %v", err)
			}
			topicName := cmd.Flag("name").Value.String()
			printMessage(c, topicName)
		},
	}
	var topicName string
	createCmd.Flags().StringVarP(&topicName, "name", "n", "", "specific topic name")
	err := createCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
	return createCmd
}

func printMessage(client sarama.Client, topic string) {
	defer client.Close()
	// 获取所有分区
	partitions, err := client.Partitions(topic)
	if err != nil {
		log.Fatalf("Failed to get partitions: %v", err)
	}

	// 遍历每个分区，读取最新消息
	for _, partition := range partitions {
		// 获取分区的最新 Offset（即下一条消息的位置）
		latestOffset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Printf("Failed to get latest offset for partition %d: %v", partition, err)
			continue
		}

		// 如果 Offset 为 0，说明分区无消息
		if latestOffset == 0 {
			fmt.Printf("Partition %d is empty\n", partition)
			continue
		}

		// 实际最后一条消息的 Offset = latestOffset - 1
		targetOffset := latestOffset - 1

		// 创建分区消费者
		consumer, err := sarama.NewConsumerFromClient(client)
		if err != nil {
			log.Printf("Failed to create consumer for partition %d: %v", partition, err)
			continue
		}
		defer consumer.Close()

		// 定位到指定 Offset
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, targetOffset)
		if err != nil {
			log.Printf("Failed to consume partition %d at offset %d: %v", partition, targetOffset, err)
			continue
		}
		defer partitionConsumer.Close()

		// 读取消息（设置超时避免阻塞）
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("  Partition %d, Offset %d:\n", partition, msg.Offset)
			fmt.Printf("  Key: %s\n", string(msg.Key))
			fmt.Printf("  Value: %s\n", string(msg.Value))
			fmt.Printf("  Timestamp: %v\n\n", msg.Timestamp)
		case <-time.After(5 * time.Second): // 5秒超时
			fmt.Printf("Timeout: No message found at offset %d (partition %d)\n\n", targetOffset, partition)
		}
	}
}
