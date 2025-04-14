package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sigs.k8s.io/yaml"
)

type Config struct {
	ApiVersion     string    `json:"apiVersion"`
	Kind           string    `json:"kind"`
	Clusters       []cluster `json:"clusters"`
	CurrentContext string    `json:"current-context"`
}

type cluster struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

var kfFile = "/Users/rui/GolandProjects/kf/kf.yaml"

func main() {

	kfCmd := &cobra.Command{
		Use:     "kafkaCm",
		Version: "v1",
		Long:    "kf is a command line tool for manage with kafka",
		Short:   "kafka is a command line tool for manage with kafka",
		Aliases: []string{"kf"},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Usage()
			if err != nil {
				return err
			}
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}

	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "kafka cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

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
			if _, err = os.Stat(kfFile); err != nil {
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
				kfConfig, err = parse2Struct(kfFile)
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
			return os.WriteFile(kfFile, yamlData, 0644)
		},
	}
	var fileName string
	addCmd.Flags().StringVarP(&fileName, "fileName", "f", "localhost:9092", "specific cluster fileName")
	err := addCmd.MarkFlagRequired("fileName")
	if err != nil {
		panic(err)
	}
	// list
	listCmd := &cobra.Command{
		Use:     "list ",
		Short:   "list kafka cluster",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err = os.Stat(kfFile); err != nil {
				//如果不存在
				if os.IsNotExist(err) {
					fmt.Println("kafka cluster does not exist")
					return nil
				}
			}
			config, err := parse2Struct(kfFile)
			for _, currentCluster := range config.Clusters {
				fmt.Println(currentCluster.Name)
			}
			return err
		},
	}
	// remove
	removeCmd := &cobra.Command{
		Use:     "remove ",
		Short:   "remove kafka cluster",
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := parse2Struct(kfFile)
			name := cmd.Flags().Lookup("name").Value.String()
			removeClusterFromName(kfFile, name, config)
			return err
		},
	}
	var clusterName string
	removeCmd.Flags().StringVarP(&clusterName, "name", "n", "", "deleting cluster name")
	err = removeCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}

	// rename
	renameCmd := &cobra.Command{
		Use:     "rename ",
		Short:   "rename kafka cluster",
		Aliases: []string{"re"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := parse2Struct(kfFile)
			name := cmd.Flags().Lookup("name").Value.String()
			renameClusterFromName(kfFile, name, config)
			return err
		},
	}

	clusterCmd.AddCommand(addCmd)
	clusterCmd.AddCommand(listCmd)
	clusterCmd.AddCommand(removeCmd)
	clusterCmd.AddCommand(renameCmd)

	kfCmd.AddCommand(clusterCmd)

	if err := kfCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}

func renameClusterFromName(file string, name string, config *Config) {

}

func removeClusterFromName(fileName, name string, config *Config) {
	// 将文件序列成对象
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}
	// 删除记录
	for i, cluster := range config.Clusters {
		if cluster.Name == name {
			config.Clusters = append(config.Clusters[:i], config.Clusters[i+1:]...)
		}
	}
	// 重置当前的上下文
	if name == config.CurrentContext {
		if len(config.Clusters) == 0 {
			config.CurrentContext = ""
		} else {
			config.CurrentContext = config.Clusters[0].Name
		}
	}

	// 重新写入yaml
	yamlData, err := yaml.Marshal(config)
	err = os.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		panic(err)
	}

}

func printCurrentClusterInfo(currentCluster cluster) {
	println(currentCluster.Name)
	println(currentCluster.Addr)
}

func parse2Struct(fileName string) (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
