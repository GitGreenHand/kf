package consumer

import (
	"github.com/spf13/cobra"
)

func NewConsumerCmd() *cobra.Command {
	consumerCmd := &cobra.Command{
		Use:     "consumer",
		Short:   "kafka consumer",
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	consumerCmd.AddCommand(NewReceiveCmd())
	return consumerCmd
}
