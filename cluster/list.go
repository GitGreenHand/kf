package cluster

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list ",
		Short:   "list kafka cluster",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(KfFile); err != nil {
				//如果不存在
				if os.IsNotExist(err) {
					fmt.Println("kafka cluster does not exist")
					return nil
				}
			}
			config, err := parse2Struct(KfFile)
			for _, currentCluster := range config.Clusters {
				fmt.Println(currentCluster.Name)
			}
			return err
		},
	}
	return listCmd
}
