package src

import (
	"fmt"
	"gitlab.cowave.com/gogo/clipboard/src/config"
	"gitlab.cowave.com/gogo/clipboard/src/jobs"
	"gitlab.cowave.com/gogo/clipboard/src/routers"
	"gitlab.cowave.com/gogo/clipboard/src/service_context"
	"gitlab.cowave.com/gogo/flaskgo"
)

const (
	Title   = "clip-sync"
	Version = "0.1.1"
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
