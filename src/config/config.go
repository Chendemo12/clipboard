package config

import (
	"gitlab.cowave.com/gogo/functools/environ"
	"sync"
)

var (
	once           = &sync.Once{}
	confInstance   *Configuration
	LocalClipboard = "" // 本机剪切板数据
)

type Configuration struct {
	HTTP struct {
		Host string `json:"host"` // API host
		Port string `json:"port"` // API port
	}
	Remote struct {
		Host         string `json:"host"`          // API host
		Port         string `json:"port"`          // API port
		SyncInterval int    `json:"sync_interval"` // 同步间隔，默认2s
	}
}

// GetConfiguration 获取配置文件
// @return *Configuration 配置文件指针
func GetConfiguration() *Configuration {
	once.Do(func() {
		confInstance = &Configuration{}
	})
	return confInstance
}

// LoadEnvirons 加载环境变量
// @param filepath string 环境变量文件路径
func (c *Configuration) LoadEnvirons(filepath string) error {
	if environ.DoesFileExists(filepath) {
		if err := environ.LoadDotenv(filepath); err != nil {
			return err
		}
	}
	c.HTTP.Host = environ.GetString("HTTP_API_HOST", "0.0.0.0")
	c.HTTP.Port = environ.GetString("HTTP_API_PORT", "7077")

	c.Remote.Host = environ.GetString("REMOTE_HTTP_HOST", "10.64.70.25")
	c.Remote.Port = environ.GetString("REMOTE_HTTP_PORT", 7078)
	c.Remote.SyncInterval = environ.GetInt("REMOTE_SYNC_INTERVAL", 2) % 10
	if c.Remote.SyncInterval <= 0 {
		c.Remote.SyncInterval = 2
	}
	return nil
}
