package plugin

import (
	"errors"
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/panel"
	"github.com/hashicorp/go-plugin"
)

type Server struct {
	c    *Config
	impl any
}

func NewServer(c *Config, impl any) (*Server, error) {
	switch impl.(type) {
	case core.Core:
	case panel.Panel:
	default:
		return nil, errors.New("unknown plugin type")
	}
	return &Server{c: c, impl: impl}, nil
}

func (s *Server) Run() error {
	hs, err := s.c.getHandshake()
	if err != nil {
		return err
	}
	pm, err := s.c.getPluginMap(s.impl)
	if err != nil {
		return err
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hs,
		Plugins:         pm,
		Logger:          s.c.Logger,
	})
	return nil
}
