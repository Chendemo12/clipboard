package src

import (
	"fmt"
	"github.com/Chendemo12/clipboard/src/config"
	"github.com/Chendemo12/clipboard/src/jobs"
	"github.com/Chendemo12/clipboard/src/routers"
	"github.com/Chendemo12/clipboard/src/service_context"
	"github.com/Chendemo12/flaskgo"
)

const (
	Title   = "clip-sync"
	Version = "0.1.2"
)

func Application(envfile string, debug bool) {
	conf := config.GetConfiguration()

	err := conf.LoadEnvirons(envfile)
	if err != nil {
		panic(err)
	}

	ctx := service_context.GetServiceContext()
	ctx.Conf = conf

	app := flaskgo.NewFlaskGo(Title, Version, debug, ctx)
	app.SetDescription("跨剪切板同步工具.").
		IncludeRouter(routers.ObjectRouter).
		RunCronjob(jobs.Cronjob)

	app.Run(conf.HTTP.Host, conf.HTTP.Port)
}

// Readme 输出版本等说明信息
func Readme() {
	fmt.Println("Application: " + Title)
	fmt.Println("Version: " + Version)
}
