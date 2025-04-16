package combo

import (
	"github.com/spf13/cobra"
	"kf/topic"
)

func NewTlsCmd() *cobra.Command {
	clusterCmd := &cobra.Command{
		Use:     "tls",
		Short:   "combo command to get topic list from cluster",
		Example: "kf tls",
		Run: func(cmd *cobra.Command, args []string) {
			topic.ListTopics()
		},
	}
	return clusterCmd
}
