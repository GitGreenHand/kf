package producer

import (
	"github.com/spf13/cobra"
)

func NewProducerCmd() *cobra.Command {
	clusterCmd := &cobra.Command{
		Use:   "producer",
		Short: "kafka producer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	clusterCmd.AddCommand(NewSendCmd())
	//clusterCmd.AddCommand(NewListCmd())
	//clusterCmd.AddCommand(NewRemoveCmd())
	//clusterCmd.AddCommand(NewSwitchCmd())
	//clusterCmd.AddCommand(NewCurrentCmd())
	//clusterCmd.AddCommand(NewRenameCmd())
	return clusterCmd
}
