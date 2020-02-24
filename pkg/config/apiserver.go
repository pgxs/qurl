package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"pgxs.io/chassis/config"
)

const (
	configEnvKey = "QURL_CONF"
)

//Server server自定义配置
type ServerConfig struct {
	Qurl QUrlConfig
}

//QUrlConfig qurl配置
type QUrlConfig struct {
	Prefix    string
	CacheSize int `yaml:"cache-size"`
}

var (
	serverConfig     ServerConfig
	serverConfigOnce sync.Once
)

//ResetEnvKey 重设配置文件环境变量名称
func ResetEnvKey() {
	fmt.Printf("Load config setting:\nEnv : %s\nFile: %s\n", configEnvKey, os.Getenv(configEnvKey))
	config.SetLoadFileEnvKey(configEnvKey)
}

//LoadServer 从配置文件加载server配置
func LoadServer() {
	serverConfigOnce.Do(func() {
		if err := config.LoadCustomFromFile(os.Getenv(config.LoadFileEnvKey()), &serverConfig); err != nil {
			log.Fatalln(err)
		}
	})
}

//Server server自定义配置
func Server() ServerConfig {
	return serverConfig
}
