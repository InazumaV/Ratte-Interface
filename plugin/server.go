package plugin

import "github.com/hashicorp/go-plugin"

type Server struct {
	c    *Config
	impl any
}

func NewServer(c *Config, impl any) *Server {
	return &Server{c: c, impl: impl}
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
