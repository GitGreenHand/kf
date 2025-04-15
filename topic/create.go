package topic

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"kf/client"
	"log"
)

func NewCreateCmd() *cobra.Command {

	createCmd := &cobra.Command{
		Use:     "create",
		Short:   "create topic to  kafka",
		Aliases: []string{"c"},
		Run: func(cmd *cobra.Command, args []string) {
			admin, err := client.GetKafkaAdmin()
			if err != nil {
				log.Fatal(err)
			}

			topicName := cmd.Flag("name").Value.String()
			err = admin.CreateTopic(topicName, &sarama.TopicDetail{
				NumPartitions:     1,
				ReplicationFactor: 1,
				ReplicaAssignment: nil,
				ConfigEntries:     nil,
			}, false)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Create topic %s successfully\n", topicName)
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
