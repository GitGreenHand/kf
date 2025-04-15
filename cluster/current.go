package cluster

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCurrentCmd() *cobra.Command {
	currentCmd := &cobra.Command{
		Use:     "current ",
		Short:   "current kafka cluster",
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := parse2Struct(KfFile)
			for _, currentCluster := range config.Clusters {
				if currentCluster.Name == config.CurrentContext {
					printCurrentClusterInfo(*currentCluster)
				}
			}
			return err
		},
	}
	return currentCmd
}

func printCurrentClusterInfo(currentCluster Cluster) {
	fmt.Println(fmt.Sprintf("name:%s", currentCluster.Name))
	fmt.Println(fmt.Sprintf("addr:%s", currentCluster.Addr))
}
