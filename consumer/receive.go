package consumer

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"kf/client"
	"log"
)

func NewReceiveCmd() *cobra.Command {
	// send
	receiveCmd := &cobra.Command{
		Use:     "receive ",
		Short:   "receive message from  topic",
		Aliases: []string{"r"},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := client.GetClient()
			if err != nil {
				return err
			}
			topic := cmd.Flags().Lookup("topic").Value.String()
			return ReceiveMessageFromTopic(client, topic)
		},
	}
	var topic string
	receiveCmd.Flags().StringVarP(&topic, "topic", "t", "", "specific message topic")
	err := receiveCmd.MarkFlagRequired("topic")
	if err != nil {
		panic(err)
	}

	return receiveCmd
}

func ReceiveMessageFromTopic(client sarama.Client, topic string) error {
	// 获取 Topic 的所有分区
	partitions, err := client.Partitions(topic)
	if err != nil {
		log.Fatal("Error getting partitions:", err)
		return err
	}
	// 遍历每个分区，消费最后一条消息
	for _, partition := range partitions {
		// 获取分区的最新 Offset
		latestOffset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Printf("Error getting latest offset for partition %d: %v", partition, err)
			continue
		}

		if latestOffset == 0 {
			fmt.Printf("Partition %d has no messages\n", partition)
			continue
		}

		// 消费 Offset = latestOffset-1 的消息（最后一条）
		consumer, err := sarama.NewConsumerFromClient(client)
		if err != nil {
			log.Printf("Error creating consumer for partition %d: %v", partition, err)
			continue
		}
		defer consumer.Close()

		pc, err := consumer.ConsumePartition(topic, partition, latestOffset-1)
		if err != nil {
			log.Printf("Error consuming partition %d: %v", partition, err)
			continue
		}
		defer pc.Close()

		select {
		case msg := <-pc.Messages():
			fmt.Printf("Last message:\n Partition %d \n Offset= %d \n Key= %s \n Value= %s\n",
				partition, msg.Offset, string(msg.Key), string(msg.Value))
		case err := <-pc.Errors():
			log.Printf("Error reading message from partition %d: %v", partition, err)
		}
	}
	return nil
}
