package producer

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"kf/client"
)

func NewSendCmd() *cobra.Command {
	// send
	sendCmd := &cobra.Command{
		Use:   "send ",
		Short: "send message to  topic",
		RunE: func(cmd *cobra.Command, args []string) error {
			producer, err := client.GetProducer()
			if err != nil {
				return err
			}
			key := cmd.Flags().Lookup("key").Value.String()
			value := cmd.Flags().Lookup("value").Value.String()
			topic := cmd.Flags().Lookup("topic").Value.String()
			return sendMessage2topic(producer, topic, key, value)
		},
	}
	var key string
	sendCmd.Flags().StringVarP(&key, "key", "k", "", "specific message key")
	var value string
	sendCmd.Flags().StringVarP(&value, "value", "v", "", "specific message value")
	err := sendCmd.MarkFlagRequired("value")
	if err != nil {
		panic(err)
	}
	var topic string
	sendCmd.Flags().StringVarP(&topic, "topic", "t", "", "specific message topic")
	err = sendCmd.MarkFlagRequired("topic")
	if err != nil {
		panic(err)
	}

	return sendCmd
}

func sendMessage2topic(c sarama.SyncProducer, topic string, key string, value string) error {
	_, _, err := c.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	})
	if err != nil {
		return err
	}
	fmt.Println("success send message to topic", topic)
	return nil
}
