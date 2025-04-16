package consumer

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"testing"
)

// 定义的接口生命当声明，= 是待实现的实体
var _ sarama.ConsumerGroupHandler = &ConsumerGroupReceiver{}

type ConsumerGroupReceiver struct {
}

func (c ConsumerGroupReceiver) Setup(session sarama.ConsumerGroupSession) error {

	return nil
}

func (c ConsumerGroupReceiver) Cleanup(session sarama.ConsumerGroupSession) error {

	return nil
}

func (c ConsumerGroupReceiver) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	msg := <-claim.Messages()
	fmt.Printf("Topic:%s Partition:%d Offset:%d \n Key:%s \n Value:%s\n",
		msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	session.MarkMessage(msg, "") // 标记消息已处理
	return nil
}

func TestConsumeGroup(t *testing.T) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Return.Errors = true
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	groupID := "kf-group"
	brokers := []string{"127.0.0.1:9092"}
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, saramaConfig)
	if err != nil {
		t.Fatalf("Error creating consumer: %v", err)
	}
	topics := []string{"test1"}
	err = consumerGroup.Consume(context.Background(), topics, &ConsumerGroupReceiver{})
	if err != nil {
		return
	}
}

func TestConsumer(t *testing.T) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建 Kafka 客户端
	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal("Error creating client:", err)
	}
	defer client.Close()

	var topic = "test1"
	// 获取 Topic 的所有分区
	partitions, err := client.Partitions(topic)
	if err != nil {
		log.Fatal("Error getting partitions:", err)
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
			fmt.Printf("Partition %d - Last message: Offset=%d Key=%s Value=%s\n",
				partition, msg.Offset, string(msg.Key), string(msg.Value))
		case err := <-pc.Errors():
			log.Printf("Error reading message from partition %d: %v", partition, err)
		}
	}
}
