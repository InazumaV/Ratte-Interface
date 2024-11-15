package plugin

import (
	"errors"
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/panel"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type Server struct {
	l  hclog.Logger
	hs plugin.HandshakeConfig
	pm map[string]plugin.Plugin
}

func NewServer(l hclog.Logger, impl any) (*Server, error) {
	var hs plugin.HandshakeConfig
	var pm map[string]plugin.Plugin
	switch impl.(type) {
	case core.Core:
		hs = core.HandshakeConfig
		pm = map[string]plugin.Plugin{
			core.PluginName: core.NewPlugin(impl.(core.Core)),
		}
	case panel.Panel:
		hs = panel.HandshakeConfig
		pm = map[string]plugin.Plugin{
			panel.PluginName: panel.NewPlugin(impl.(panel.Panel)),
		}
	default:
		return nil, errors.New("the plugin type is not supported")
	}
	return &Server{l: l, hs: hs, pm: pm}, nil
}

func (s *Server) Run() error {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: s.hs,
		Plugins:         s.pm,
		Logger:          s.l,
	})
	return nil
}
