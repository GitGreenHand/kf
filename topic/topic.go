package topic

import "github.com/spf13/cobra"

func NewTopicCmd() *cobra.Command {

	topicCmd := &cobra.Command{
		Use:     "topic",
		Short:   "kafka topic",
		Aliases: []string{"t"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	topicCmd.AddCommand(NewListCmd())
	topicCmd.AddCommand(NewCreateCmd())
	topicCmd.AddCommand(NewDeleteCmd())
	topicCmd.AddCommand(NewDescribeCmd())
	topicCmd.AddCommand(NewMessageCmd())

	return topicCmd

}
