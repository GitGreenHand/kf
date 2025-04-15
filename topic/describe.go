package topic

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"kf/client"
	"log"
)

func NewDescribeCmd() *cobra.Command {

	createCmd := &cobra.Command{
		Use:     "describe",
		Short:   "describe topic to  kafka",
		Aliases: []string{"ds"},
		Run: func(cmd *cobra.Command, args []string) {
			client, err := client.GetClient()
			if err != nil {
				log.Fatal(err)
			}
			topicName := cmd.Flag("name").Value.String()
			describeTopic(client, topicName)
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

func describeTopic(client sarama.Client, topic string) {

	// 获取 Topic 的所有分区
	partitions, err := client.Partitions(topic)
	if err != nil {
		log.Fatalf("Error getting partitions: %v", err)
	}
	totalMessages := int64(0)

	// 遍历每个分区
	for _, partition := range partitions {
		// 获取分区的 Oldest Offset
		oldestOffset, err := client.GetOffset(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Printf("Error getting oldest offset for partition %d: %v", partition, err)
			continue
		}

		// 获取分区的 Latest Offset
		latestOffset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Printf("Error getting latest offset for partition %d: %v", partition, err)
			continue
		}

		// 计算当前分区的消息数量
		messagesInPartition := latestOffset - oldestOffset
		totalMessages += messagesInPartition

		fmt.Printf("  Partition %d:\n", partition)
		fmt.Printf("  Oldest Offset: %d\n", oldestOffset)
		fmt.Printf("  Latest Offset: %d\n", latestOffset)
		fmt.Printf("  Messages: %d\n\n", messagesInPartition)
	}

	fmt.Printf("Total messages in topic %s: %d\n", topic, totalMessages)
}
