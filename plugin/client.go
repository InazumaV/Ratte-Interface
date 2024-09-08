package plugin

import (
	"errors"
	"github.com/hashicorp/go-plugin"
)

type Client[T any] struct {
	c   *plugin.Client
	obj T
}

func NewClient[T any](c *Config) (client *Client[T], err error) {
	hs, err := c.getHandshake()
	if err != nil {
		return nil, err
	}
	pm, err := c.getPluginMap(nil)
	if err != nil {
		return nil, err
	}
	cli := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: hs,
		Plugins:         pm,
		Cmd:             c.Cmd,
		Logger:          c.Logger,
	})
	rpc, err := cli.Client()
	if err != nil {
		return client, err
	}
	v, err := rpc.Dispense(pluginName)
	if err != nil {
		return client, err
	}
	switch v.(type) {
	case T:
		client.obj = v.(T)
	default:
		return client, errors.New("unknown plugin type")
	}
	return
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
