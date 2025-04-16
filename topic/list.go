package topic

import (
	"fmt"
	"github.com/spf13/cobra"
	"kf/client"
	"log"
)

func NewListCmd() *cobra.Command {

	listCmd := &cobra.Command{
		Use:     "list",
		Short:   "list topic",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			ListTopics()
		},
	}
	return listCmd
}

func ListTopics() {
	admin, err := client.GetKafkaAdmin()
	if err != nil {
		log.Fatal(err)
	}
	topics, err := admin.ListTopics()
	if err != nil {
		log.Fatal(err)
	}
	for topic := range topics {
		fmt.Println(topic)
	}
}
