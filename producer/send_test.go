package producer

import (
	"github.com/IBM/sarama"
	"testing"
)

func TestSend(t *testing.T) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true          // 必须为 true
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll // 确保消息可靠写入
	saramaConfig.Producer.Retry.Max = 5                    // 重试次数
	brokers := []string{"127.0.0.1:9092"}
	producer, err := sarama.NewSyncProducer(brokers, saramaConfig)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: "test1",
		Value: sarama.StringEncoder("hello world"),
		Key:   sarama.StringEncoder("key"),
	})
	if err != nil {
		t.Fatal(err)
	}

}
