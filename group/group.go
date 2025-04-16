package group

import "github.com/spf13/cobra"

func NewGroupCmd() *cobra.Command {
	groupCmd := &cobra.Command{
		Use:   "group",
		Short: "kafka consumer group",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	groupCmd.AddCommand(NewListCmd())
	return groupCmd
}
