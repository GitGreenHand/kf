package client

import (
	"errors"
	"github.com/IBM/sarama"
	"kf/cluster"
	"log"
	"os"
	"sigs.k8s.io/yaml"
	"strings"
)

// GetKafkaAdmin  todo 后续使用channel来进行优化
func GetKafkaAdmin() (sarama.ClusterAdmin, error) {
	// 获取当期的配置文件,创建client
	config, err := getCurrentClusterConfig()
	if err != nil {
		return nil, err
	}
	saramaConfig := sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin(splitKafkaAddr(config.Addr), saramaConfig)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// GetClient  获取client
func GetClient() (sarama.Client, error) {
	// 获取当期的配置文件,创建client
	config, err := getCurrentClusterConfig()
	if err != nil {
		return nil, err
	}
	saramaConfig := sarama.NewConfig()
	consumer, err := sarama.NewClient(splitKafkaAddr(config.Addr), saramaConfig)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	return consumer, nil
}

// 使用;来分割
func splitKafkaAddr(addr string) []string {
	return strings.Split(addr, ";")
}

func getCurrentClusterConfig() (*cluster.Cluster, error) {
	config := &cluster.Config{}
	file, err := os.ReadFile(cluster.KfFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	if config.CurrentContext == "" || len(config.Clusters) == 0 {
		return nil, errors.New("no current cluster found,please add cluster context")
	}
	for _, _cluster := range config.Clusters {
		if config.CurrentContext == _cluster.Name {
			return _cluster, nil
		}
	}
	return nil, errors.New("cluster not found")
}
