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
		Short:   "kafka is a command line tool for manage with kafka",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
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
		},
	}
	return listCmd
}
