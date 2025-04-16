package combo

import (
	"github.com/spf13/cobra"
	"kf/client"
	pro "kf/producer"
)

func NewPseCmd() *cobra.Command {
	pseCmd := &cobra.Command{
		Use:     "pse",
		Short:   "combo command to send one message to given topic",
		Example: "kf pse -n topic message",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			producer, err := client.GetProducer()
			if err != nil {
				return err
			}
			topic := cmd.Flags().Lookup("topic").Value.String()
			value := args[0]
			key := ""
			return pro.SendMessage2topic(producer, topic, key, value)
		},
	}
	var topic string
	pseCmd.Flags().StringVarP(&topic, "topic", "t", "", "specific message topic")
	err := pseCmd.MarkFlagRequired("topic")
	if err != nil {
		panic(err)
	}
	return pseCmd
}
