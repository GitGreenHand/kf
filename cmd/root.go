package cmd

import (
	"github.com/spf13/cobra"
	"kf/cluster"
	"kf/combo"
	"kf/consumer"
	"kf/group"
	"kf/producer"
	"kf/topic"
)

func NewRootCmd() *cobra.Command {
	kfCmd := &cobra.Command{
		Use:     "kafkaCm",
		Version: "v1.0",
		Long:    "kf is a command line tool for manage with kafka",
		Short:   "kafka is a command line tool for manage with kafka",
		Aliases: []string{"kf"},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Usage()
			if err != nil {
				return err
			}
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}
	// 添加子命令
	kfCmd.AddCommand(cluster.NewClusterCmd())
	kfCmd.AddCommand(topic.NewTopicCmd())
	kfCmd.AddCommand(producer.NewProducerCmd())
	kfCmd.AddCommand(group.NewGroupCmd())
	kfCmd.AddCommand(consumer.NewConsumerCmd())
	// combo 命令
	kfCmd.AddCommand(combo.NewTlsCmd())
	kfCmd.AddCommand(combo.NewPseCmd())
	kfCmd.AddCommand(combo.NewCreCmd())
	return kfCmd
}
