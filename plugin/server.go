package plugin

import (
	"errors"
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/panel"
	"github.com/hashicorp/go-plugin"
)

type Server struct {
	c  *Config
	pm map[string]plugin.Plugin
}

func NewServer(c *Config, impl any) (*Server, error) {
	var pm map[string]plugin.Plugin
	switch impl.(type) {
	case core.Core:
		pm = map[string]plugin.Plugin{
			core.PluginName: core.NewPlugin(impl.(core.Core)),
		}
	case panel.Panel:
		pm = map[string]plugin.Plugin{
			panel.PluginName: panel.NewPlugin(impl.(panel.Panel)),
		}
	default:
		return nil, errors.New("unknown plugin type")
	}
	return &Server{c: c, pm: pm}, nil
}

func (s *Server) Run() error {
	hs, err := s.c.getHandshake()
	if err != nil {
		return err
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hs,
		Plugins:         s.pm,
		Logger:          s.c.Logger,
	})
	return nil
}
