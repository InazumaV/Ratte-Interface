package baseplugin

import (
	"errors"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"os/exec"
)

type Client struct {
	c *plugin.Client
}

func NewClient(
	name string,
	cmd *exec.Cmd,
	l hclog.Logger,
	p plugin.Plugin,
	hs plugin.HandshakeConfig) (client *Client, obj any, err error) {
	cli := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: hs,
		Plugins: map[string]plugin.Plugin{
			name: p,
		},
		Cmd:    cmd,
		Logger: l,
	})
	rpc, err := cli.Client()
	if err != nil {
		return nil, nil, err
	}
	v, err := rpc.Dispense(name)
	if err != nil {
		return nil, nil, err
	}
	if v == nil {
		return nil, nil, errors.New("get object failed, please check log")
	}
	return &Client{
		c: cli,
	}, v, nil
}

func (c *Client) RawClient() *plugin.Client {
	return c.c
}

func (c *Client) Close() {
	c.c.Kill()
}
