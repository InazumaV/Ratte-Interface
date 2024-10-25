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
	CoreType  = "core"
	PanelType = "node"
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
