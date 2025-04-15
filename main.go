package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sigs.k8s.io/yaml"
)

type Config struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	// 如果是值类型，无法修改cluster中的值，如果需要修改，需要定义成指针类型
	Clusters       []*cluster `json:"clusters"`
	CurrentContext string     `json:"current-context"`
}

type cluster struct {
	ID   string `json:"id,omitempty"`
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
			oldName := cmd.Flags().Lookup("oldName").Value.String()
			newName := cmd.Flags().Lookup("newName").Value.String()
			renameClusterFromName(oldName, newName, config)
			return err
		},
	}
	var oldName string
	var newName string
	renameCmd.Flags().StringVarP(&oldName, "oldName", "o", "", "switch cluster oldName")
	err = renameCmd.MarkFlagRequired("oldName")
	if err != nil {
		panic(err)
	}
	renameCmd.Flags().StringVarP(&newName, "newName", "n", "", "switch cluster newName")
	err = renameCmd.MarkFlagRequired("newName")
	if err != nil {
		panic(err)
	}

	// switch
	switchCmd := &cobra.Command{
		Use:     "switch ",
		Short:   "switch kafka cluster",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := parse2Struct(kfFile)
			name := cmd.Flags().Lookup("name").Value.String()
			switchClusterFromName(kfFile, name, config)
			return err
		},
	}
	currentCmd := &cobra.Command{
		Use:     "current ",
		Short:   "current kafka cluster",
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := parse2Struct(kfFile)
			for _, currentCluster := range config.Clusters {
				if currentCluster.Name == config.CurrentContext {
					printCurrentClusterInfo(*currentCluster)
				}
			}
			return err
		},
	}
	var switchClusterName string
	switchCmd.Flags().StringVarP(&switchClusterName, "name", "n", "", "switch cluster name")
	err = switchCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}

	clusterCmd.AddCommand(addCmd)
	clusterCmd.AddCommand(listCmd)
	clusterCmd.AddCommand(removeCmd)
	clusterCmd.AddCommand(switchCmd)
	clusterCmd.AddCommand(currentCmd)
	clusterCmd.AddCommand(renameCmd)

	kfCmd.AddCommand(clusterCmd)

	if err := kfCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}

func switchClusterFromName(fileName string, name string, config *Config) {
	// 将文件序列成对象
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}
	// 替换当前的上下文
	config.CurrentContext = name
	// 重新写入yaml
	yamlData, err := yaml.Marshal(config)
	err = os.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		panic(err)
	}

}

func renameClusterFromName(oldName, newName string, config *Config) {

	// 更新current-context
	if config.CurrentContext == oldName {
		config.CurrentContext = newName
	}

	// 更新名称
	for _, cluster := range config.Clusters {
		if oldName == cluster.Name {
			cluster.Name = newName
			println(cluster.Name)
		}
	}
	// 回写
	yamlData, err := yaml.Marshal(config)
	fmt.Println(string(yamlData))
	err = os.WriteFile(kfFile, yamlData, 0644)
	if err != nil {
		panic(err)
	}

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
	fmt.Println(fmt.Sprintf("name:%s", currentCluster.Name))
	fmt.Println(fmt.Sprintf("addr:%s", currentCluster.Addr))
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
