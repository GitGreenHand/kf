package group

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"kf/client"
)

func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list ",
		Short:   "list consumer  from kafka cluster",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			admin, err := client.GetKafkaAdmin()
			if err != nil {
				return err
			}
			return listConsumerGroup(admin)
		},
	}
	return listCmd
}

func listConsumerGroup(admin sarama.ClusterAdmin) error {

	groups, err := admin.ListConsumerGroups()
	if err != nil {
		return err
	}
	for groupId, _ := range groups {
		fmt.Println(groupId)
	}
	return nil
}
