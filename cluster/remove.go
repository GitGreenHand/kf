package cluster

import "github.com/spf13/cobra"

func NewRemoveCmd() *cobra.Command {
	removeCmd := &cobra.Command{
		Use:     "remove ",
		Short:   "remove kafka cluster",
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := parse2Struct(KfFile)
			name := cmd.Flags().Lookup("name").Value.String()
			removeClusterFromName(KfFile, name, config)
			return err
		},
	}
	var clusterName string
	removeCmd.Flags().StringVarP(&clusterName, "name", "n", "", "deleting cluster name")
	err := removeCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
	return removeCmd
}
