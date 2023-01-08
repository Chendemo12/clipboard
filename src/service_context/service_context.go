package service_context

import (
	"context"
	"github.com/Chendemo12/clipboard/src/config"
	"sync"
)

var (
	once = &sync.Once{}
	sc   *ServiceContext
)

type ServiceContext struct {
	Conf    *config.Configuration
	Context *context.Context
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
