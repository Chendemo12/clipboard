package client

import (
	"gitlab.cowave.com/gogo/functools/httpc"
	"gitlab.cowave.com/gogo/functools/zaplog"
)

type Clipboard struct {
	Host    string       `json:"host"`
	Port    string       `json:"port"`
	Timeout int          `json:"timeout"`
	Logger  zaplog.Iface `json:"logger"`
	client  *httpc.Httpr
	i       bool
}

func (c *Clipboard) init() *Clipboard {
	if !c.i {
		c.client = &httpc.Httpr{Host: c.Host, Port: c.Port}
		c.client.SetTimeout(5).SetLogger(c.Logger).NewClient()
	}
	c.i = true
	return c
}

func (c *Clipboard) GetText() (string, error) {
	c.init()
	resp, err := c.client.Get("/clipboard", map[string]string{})
	if err != nil {
		return "", err
	}
	text := ""
	err = c.client.UnmarshalJson(resp, &text)
	if err != nil {
		return "", err
	}
	return text, nil
}

func (c *Clipboard) SetText(text string) error {
	c.init()
	_, err := c.client.Post("/clipboard", map[string]any{"message": text})
	return err
}
