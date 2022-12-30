package service_context

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gitlab.cowave.com/gogo/clipboard/src/config"
	"gitlab.cowave.com/gogo/functools/kafkac"
	"gorm.io/gorm"
	"sync"
)

var (
	once = &sync.Once{}
	sc   *ServiceContext
)

type ServiceContext struct {
	Conf         *config.Configuration
	DB           *gorm.DB
	Kafka        *kafkac.KafkaClient
	RedisBFF     *redis.Client
	RedisDefault *redis.Client
	Context      *context.Context
}

func (c *ServiceContext) Config() any { return c.Conf }

// GetServiceContext 获取 service 信息
// @return *ServiceContext 配置文件指针
func GetServiceContext() *ServiceContext {
	once.Do(func() {
		sc = &ServiceContext{}
	})
	return sc
}
