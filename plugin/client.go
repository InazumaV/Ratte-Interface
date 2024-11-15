package plugin

import (
	"errors"
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/panel"
	"github.com/hashicorp/go-plugin"
)

type Client struct {
	c   *plugin.Client
	obj any
}

const (
	TypeCore  = 0
	TypePanel = 1
)

func NewClient(Type int, c *Config) (client *Client, err error) {
	var name string
	var pm = map[string]plugin.Plugin{}
	var hs = plugin.HandshakeConfig{}
	switch Type {
	case TypeCore:
		name = core.PluginName
		hs = core.HandshakeConfig
		pm = map[string]plugin.Plugin{
			core.PluginName: core.NewPlugin(nil),
		}
	case TypePanel:
		name = panel.PluginName
		hs = panel.HandshakeConfig
		pm = map[string]plugin.Plugin{
			panel.PluginName: panel.NewPlugin(nil),
		}
	default:
		return nil, errors.New("the plugin type is not supported")
	}
	cli := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: hs,
		Plugins:         pm,
		Cmd:             c.Cmd,
		Logger:          c.Logger,
	})
	rpc, err := cli.Client()
	if err != nil {
		return nil, err
	}
	v, err := rpc.Dispense(name)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, errors.New("get object failed, please check log")
	}
	return &Client{
		c:   cli,
		obj: v,
	}, nil
}

func (c *Client) RawClient() *plugin.Client {
	return c.c
}

func (c *Client) Caller() any {
	return c.obj
}

func (c *Client) Close() {
	c.c.Kill()
}
