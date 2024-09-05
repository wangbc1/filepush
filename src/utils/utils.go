package utils

import (
	"github.com/spf13/viper"
	"log"
)

// Conf 解析配置
type Conf struct {
	Remoter []Remote `yaml:"remote"`
	Locals  []Local  `yaml:"local"`
}

type Remote struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Path   string `yaml:"path"`
}

type Local struct {
	Path string `yaml:"path"`
}

// ConfigData 配置数据变量
var ConfigData *Conf

// 引入包后，初始化配置
func init() {
	// 导入配置文件
	// 定义配置文件路径
	var FilePath = "../config/application.yaml"
	// 配置文件类型为yaml
	viper.SetConfigType("yaml")
	// 配置文件路径
	viper.SetConfigFile(FilePath)

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件报错，配置文件路径为: %s", FilePath, err.Error())
	}
	err = viper.Unmarshal(&ConfigData)
	if err != nil {
		log.Fatalln("导出数据有误", err.Error())
	}
}

// GetConfigData 返回配置数据方法
func GetConfigData() *Conf {
	return ConfigData
}
