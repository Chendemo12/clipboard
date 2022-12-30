package main

import (
	"flag"
	"gitlab.cowave.com/gogo/clipboard/src"
	_ "go.uber.org/automaxprocs"
	"math/rand"
	"time"
)

var (
	debug, v, version bool
	envfile           string // 环境变量文件路径
)

func init() {
	flag.BoolVar(&v, "v", false, "show version.")
	flag.BoolVar(&version, "version", false, "show version and description.")
	flag.BoolVar(&debug, "dev", false, "run with DevMode.")
	flag.StringVar(&envfile, "env", "", "specifying a environment file.")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if v || version {
		src.Readme()
		return
	}
	if envfile == "" { // 环境变量文件
		envfile = "/etc/opt/cowave/" + src.Title + "/project.env"
	}

	src.Application(envfile, debug)
}
