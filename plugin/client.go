package plugin

import (
	"errors"
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/panel"
	"github.com/hashicorp/go-plugin"
)

type Client[T any] struct {
	c   *plugin.Client
	obj T
}

func NewClient[T any](c *Config) (client *Client[T], err error) {
	client = new(Client[T])
	hs, err := c.getHandshake()
	if err != nil {
		return nil, err
	}
	name := ""
	switch c.Type {
	case CoreType:
		name = core.PluginName
	case PanelType:
		name = panel.PluginName
	default:
		return nil, errors.New("the plugin type is not supported")
	}
	cli := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: hs,
		Plugins:         map[string]plugin.Plugin{},
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
	switch v.(type) {
	case T:
		client.obj = v.(T)
	default:
		return nil, errors.New("unknown plugin type")
	}
	return client, nil
}

func (c *Client[T]) RawClient() *plugin.Client {
	return c.c
}

func (c *Client[T]) Caller() T {
	return c.obj
}

func (c *Client[T]) Close() {
	c.c.Kill()
}
