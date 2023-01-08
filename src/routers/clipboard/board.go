package clipboard

import (
	"github.com/Chendemo12/clipboard/src/config"
	"github.com/Chendemo12/clipboard/src/utils/clipboard"
	"github.com/Chendemo12/flaskgo"
)

func ReadClipboard(c *flaskgo.Context) any {
	content, err := clipboard.ReadAll()
	if err != nil {
		content = err.Error()
	}

	return content
}

type Text struct {
	flaskgo.BaseModel
	Message string `json:"message" description:"写入剪切板的字符串"`
}

func (t Text) Doc__() string { return "写入剪切板的数据" }

func WriteClipboard(c *flaskgo.Context) any {
	content := &Text{}
	if err := c.ShouldBindJSON(content); err != nil {
		return err
	}

	err := clipboard.WriteAll(content.Message)
	if err != nil {
		return err.Error()
	}

	// 更新本机剪切板数据
	config.LocalClipboard = content.Message

	return content
}
