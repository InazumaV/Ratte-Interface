package baseplugin

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type Server struct {
	l  hclog.Logger
	hs plugin.HandshakeConfig
	pm map[string]plugin.Plugin
}

func NewServer(
	name string,
	l hclog.Logger,
	hs plugin.HandshakeConfig,
	impl plugin.Plugin) (*Server, error) {
	return &Server{l: l, hs: hs, pm: map[string]plugin.Plugin{
		name: impl,
	}}, nil
}

func (s *Server) Run() error {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: s.hs,
		Plugins:         s.pm,
		Logger:          s.l,
	})
	return nil
}
