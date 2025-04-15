package cluster

import "github.com/spf13/cobra"

func NewSwitchCmd() *cobra.Command {
	switchCmd := &cobra.Command{
		Use:     "switch ",
		Short:   "switch kafka cluster",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := parse2Struct(KfFile)
			name := cmd.Flags().Lookup("name").Value.String()
			switchClusterFromName(KfFile, name, config)
			return err
		},
	}

	var switchClusterName string
	switchCmd.Flags().StringVarP(&switchClusterName, "name", "n", "", "switch cluster name")
	err := switchCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
	return switchCmd

}
