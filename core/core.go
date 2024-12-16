package core

import (
	"github.com/InazumaV/Ratte-Interface/baseplugin"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"net/rpc"
	"os/exec"
)

const PluginName = "ratte-core"

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "is-ratte-core",
	MagicCookieValue: "true",
}

type Core interface {
	Start(dataPath string, config []byte) error
	Close() error
	AddNode(params *AddNodeParams) error
	DelNode(name string) error
	AddUsers(p *AddUsersParams) error
	GetUserTraffic(p *GetUserTrafficParams) *GetUserTrafficResponse
	ResetUserTraffic(p *ResetUserTrafficParams) error
	DelUsers(params *DelUsersParams) error
	Protocols() []string
	Type() string
}

type PluginClient struct {
	Core
	*baseplugin.Client
}

func (c *PluginClient) Close() error {
	defer c.Client.Close()
	return c.Core.Close()
}

func NewClient(l hclog.Logger, cmd *exec.Cmd) (client *PluginClient, err error) {
	pc, obj, err := baseplugin.NewClient(PluginName, cmd, l, NewPlugin(nil), HandshakeConfig)
	if err != nil {
		return nil, err
	}
	return &PluginClient{
		Client: pc,
		Core:   obj.(Core),
	}, nil
}

type PluginServer struct {
	*baseplugin.Server
}

func NewServer(l hclog.Logger, c Core) (*PluginServer, error) {
	s, err := baseplugin.NewServer(PluginName, l, HandshakeConfig, NewPlugin(c))
	if err != nil {
		return nil, err
	}
	return &PluginServer{Server: s}, nil
}

type Plugin struct {
	c Core
}

func NewPlugin(c Core) *Plugin {
	return &Plugin{c: c}
}

func (p *Plugin) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return &PluginImplServer{core: p.c}, nil
}

func (_ *Plugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginImplClient{c: c}, nil
}

type PluginImplServer struct {
	core Core
}

var _ Core = (*PluginImplClient)(nil)

type PluginImplClient struct {
	c *rpc.Client
}

func (c *PluginImplClient) call(method string, args interface{}, reply interface{}) error {
	return c.c.Call("Plugin."+method, args, reply)
}

type StartParams struct {
	DataPath string
	Config   []byte
}

func (s *PluginImplServer) Start(p *StartParams, err *error) error {
	*err = s.core.Start(p.DataPath, p.Config)
	return nil
}
func (c *PluginImplClient) Start(dataPath string, config []byte) (err error) {
	err2 := c.call("Start", &StartParams{
		DataPath: dataPath,
		Config:   config,
	}, &err)
	if err2 != nil {
		return err
	}
	return nil
}

func (s *PluginImplServer) Close(_ interface{}, err *error) error {
	*err = s.core.Close()
	return nil
}
func (c *PluginImplClient) Close() (err error) {
	err2 := c.call("Close", new(interface{}), &err)
	if err2 != nil {
		return err
	}
	return nil
}

func (s *PluginImplServer) Protocols(_ interface{}, rsp *[]string) error {
	*rsp = s.core.Protocols()
	return nil
}
func (c *PluginImplClient) Protocols() (ps []string) {
	_ = c.call("Protocols", new(interface{}), &ps)
	return ps
}

func (s *PluginImplServer) Type(_ interface{}, t *string) error {
	*t = s.core.Type()
	return nil
}
func (c *PluginImplClient) Type() (t string) {
	_ = c.call("Type", new(interface{}), &t)
	return t
}
