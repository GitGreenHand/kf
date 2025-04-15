package cluster

import (
	"github.com/spf13/cobra"
	"os"
	"sigs.k8s.io/yaml"
)

func NewAddCmd() *cobra.Command {
	// add
	addCmd := &cobra.Command{
		Use:   "add ",
		Short: "add kafka cluster",
		RunE: func(cmd *cobra.Command, args []string) error {

			var kfConfig *Config
			var config *Config
			var err error
			fileName := cmd.Flags().Lookup("fileName").Value.String()
			// judge kf file 是否存在
			if _, err = os.Stat(KfFile); err != nil {
				//如果不存在
				if os.IsNotExist(err) {
					config, err = parse2Struct(fileName)
					if err != nil {
						return err
					}
					kfConfig = config
					kfConfig.CurrentContext = config.Clusters[0].Name
				}
			} else { // 如果存在
				kfConfig, err = parse2Struct(KfFile)
				if err != nil {
					return err
				}
				config, err = parse2Struct(fileName)
				kfConfig.Clusters = append(kfConfig.Clusters, config.Clusters...)
				if kfConfig.CurrentContext == "" {
					kfConfig.CurrentContext = config.Clusters[0].Name
				}
			}
			yamlData, err := yaml.Marshal(kfConfig)
			if err != nil {
				return err
			}
			return os.WriteFile(KfFile, yamlData, 0644)
		},
	}
	var fileName string
	addCmd.Flags().StringVarP(&fileName, "fileName", "f", "localhost:9092", "specific cluster fileName")
	err := addCmd.MarkFlagRequired("fileName")
	if err != nil {
		panic(err)
	}
	return addCmd
}
