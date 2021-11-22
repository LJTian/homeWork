package conFig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 日志文件信息
type LogStruct struct {
	Path string `yaml:"Path"`
	Name string	`yaml:"Name"`
	Level int `yaml:"Level"`
}

// 网络相关信息
type NetStruct struct {
	IP string	`yaml:"IP"`
	Port string	`yaml:"Port"`
}

// 描述域
type SpecStruct struct {
	NetInfo	NetStruct `yaml:"netInfo"`
	LogInfo LogStruct `yaml:"logInfo"`
}

// 配置文件
type Config struct {
	ApiVersion 	string `yaml:"apiVersion"`
	Kind		string `yaml:"kind"`
	Metadata 	string `yaml:"metadata"`
	Spec		SpecStruct 	`yaml:"spec"`
}

// 获取配置文件内容
func GetConfig(pathStr string)(setting Config, err error){

	config, err := ioutil.ReadFile(pathStr)
	if err != nil {
		return
	}
	yaml.Unmarshal(config,&setting)
	return
}
