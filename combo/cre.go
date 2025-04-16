package combo

import (
	"github.com/spf13/cobra"
	"kf/client"
	"kf/consumer"
)

func NewCreCmd() *cobra.Command {
	creCmd := &cobra.Command{
		Use:     "cre",
		Short:   "combo command to get  newest message from given topic",
		Example: "cre topic",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := client.GetClient()
			if err != nil {
				return err
			}
			topic := args[0]
			return consumer.ReceiveMessageFromTopic(client, topic)
		},
	}
	return creCmd
}
