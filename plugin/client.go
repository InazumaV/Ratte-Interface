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
	client = new(Client[T])
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
		return nil, err
	}
	v, err := rpc.Dispense(pluginName)
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
