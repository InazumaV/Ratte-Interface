package plugin

import (
	"errors"
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/panel"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"os/exec"
)

const (
	pluginName = "plugin"
	CoreType   = "core"
	PanelType  = "node"
)

type Config struct {
	Type   string
	Cmd    *exec.Cmd
	Logger hclog.Logger
}

func (c *Config) getHandshake() (plugin.HandshakeConfig, error) {
	switch c.Type {
	case CoreType:
		return core.HandshakeConfig, nil
	case PanelType:
		return panel.HandshakeConfig, nil
	default:
		return plugin.HandshakeConfig{}, errors.New("the plugin type is not supported")
	}
}

func (c *Config) getPluginMap(impl any) (map[string]plugin.Plugin, error) {
	switch c.Type {
	case CoreType:
		c := &core.Plugin{}
		if impl != nil {
			c = core.NewPlugin(impl.(core.Core))
		}
		return map[string]plugin.Plugin{
			core.PluginName: c,
		}, nil
	case PanelType:
		p := &panel.Plugin{}
		if impl != nil {
			p = panel.NewPlugin(impl.(panel.Panel))
		}
		return map[string]plugin.Plugin{
			core.PluginName: p,
		}, nil
	default:
		return nil, errors.New("the plugin type is not supported")
	}
}
