package cluster

import (
	"os"
)

var KfFile = ""

func init() {
	wokDir := os.Getenv("WOK_DIR")
	if len(wokDir) == 0 {
		dir, _ := os.Getwd()
		wokDir = dir + "/kf.yaml"
	}
	KfFile = wokDir
}

type Config struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	// 如果是值类型，无法修改cluster中的值，如果需要修改，需要定义成指针类型
	Clusters       []*Cluster `json:"clusters"`
	CurrentContext string     `json:"current-context"`
}
