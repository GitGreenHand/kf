package cmd

import (
	"github.com/spf13/cobra"
	"kf/cluster"
	"kf/topic"
)

func NewRootCmd() *cobra.Command {
	kfCmd := &cobra.Command{
		Use:     "kafkaCm",
		Version: "v1",
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
	return kfCmd
}
