package routers

import (
	"gitlab.cowave.com/gogo/clipboard/src/routers/clipboard"
	"gitlab.cowave.com/gogo/flaskgo"
)

var ObjectRouter = flaskgo.APIRouter("/api", []string{"User"})

func init() {
	ObjectRouter.GET("/clipboard", flaskgo.String, "获取剪切板数据", clipboard.ReadClipboard)
	ObjectRouter.POST(
		"/clipboard", clipboard.Text{}, clipboard.Text{}, "写入数据到剪切板", clipboard.WriteClipboard,
	)
}
