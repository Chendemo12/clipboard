package jobs

import (
	"gitlab.cowave.com/gogo/clipboard/src/config"
	"gitlab.cowave.com/gogo/clipboard/src/utils/client"
	"gitlab.cowave.com/gogo/clipboard/src/utils/clipboard"
	"gitlab.cowave.com/gogo/flaskgo"
	"time"
)

func Cronjob(ctx *flaskgo.Service) error {
	go SyncClipboard(ctx)

	return nil
}

// SyncClipboard 当本机剪切板发生变化时，将本机数据发送到远程
func SyncClipboard(ctx *flaskgo.Service) {
	conf := config.GetConfiguration()
	clip := &client.Clipboard{
		Host:    conf.Remote.Host,
		Port:    conf.Remote.Port,
		Timeout: 2,
		Logger:  ctx.Logger(),
	}
	for {
		time.Sleep(time.Second)

		c, err := clipboard.ReadAll() // 读取本机剪切板
		if err != nil {
			ctx.Logger().Warn("read local clipboard failed, ", err.Error())
			continue
		}

		if c != config.LocalClipboard { // 本机剪切板发生变化
			// 将本机剪切板发送到远程
			if err = clip.SetText(c); err != nil {
				ctx.Logger().Warn("send clipboard failed, ", err.Error())
				continue
			}
			config.LocalClipboard = c // 剪切板发送成功，更新本地记录
		}
	}
}
