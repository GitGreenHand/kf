package cluster

import (
	"github.com/spf13/cobra"
	"os"
	"sigs.k8s.io/yaml"
)

type Cluster struct {
	//ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func NewClusterCmd() *cobra.Command {
	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "kafka cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	clusterCmd.AddCommand(NewAddCmd())
	clusterCmd.AddCommand(NewListCmd())
	clusterCmd.AddCommand(NewRemoveCmd())
	clusterCmd.AddCommand(NewSwitchCmd())
	clusterCmd.AddCommand(NewCurrentCmd())
	clusterCmd.AddCommand(NewRenameCmd())
	return clusterCmd
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
	err = os.WriteFile(KfFile, yamlData, 0644)
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
