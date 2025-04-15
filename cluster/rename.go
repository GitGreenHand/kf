package cluster

import "github.com/spf13/cobra"

func NewRenameCmd() *cobra.Command {

	renameCmd := &cobra.Command{
		Use:     "rename ",
		Short:   "rename kafka cluster",
		Aliases: []string{"re"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := parse2Struct(KfFile)
			oldName := cmd.Flags().Lookup("oldName").Value.String()
			newName := cmd.Flags().Lookup("newName").Value.String()
			renameClusterFromName(oldName, newName, config)
			return err
		},
	}
	var oldName string
	var newName string
	renameCmd.Flags().StringVarP(&oldName, "oldName", "o", "", "switch cluster oldName")
	err := renameCmd.MarkFlagRequired("oldName")
	if err != nil {
		panic(err)
	}
	renameCmd.Flags().StringVarP(&newName, "newName", "n", "", "switch cluster newName")
	err = renameCmd.MarkFlagRequired("newName")
	if err != nil {
		panic(err)
	}
	return renameCmd
}
