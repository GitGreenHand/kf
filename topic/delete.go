package topic

import (
	"fmt"
	"github.com/spf13/cobra"
	"kf/client"
	"log"
)

func NewDeleteCmd() *cobra.Command {

	createCmd := &cobra.Command{
		Use:     "delete",
		Short:   "delete topic to  kafka",
		Aliases: []string{"d"},
		Run: func(cmd *cobra.Command, args []string) {
			admin, err := client.GetKafkaAdmin()
			if err != nil {
				log.Fatal(err)
			}

			topicName := cmd.Flag("name").Value.String()
			err = admin.DeleteTopic(topicName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Delete topic %s successfully\n", topicName)
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
