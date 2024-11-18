package core

import (
	"github.com/Yuzuki616/Ratte-Interface/baseplugin"
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

type Client struct {
	Core
	*baseplugin.Client
}

func (c *Client) Close() error {
	defer c.Client.Close()
	return c.Core.Close()
}

func NewClient(l hclog.Logger, cmd *exec.Cmd) (client *Client, err error) {
	pc, obj, err := baseplugin.NewClient(PluginName, cmd, l, NewPlugin(nil), HandshakeConfig)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: pc,
		Core:   obj.(Core),
	}, nil
}

type Server struct {
	*baseplugin.Server
}

func NewServer(l hclog.Logger, c Core) (*Server, error) {
	s, err := baseplugin.NewServer(PluginName, l, HandshakeConfig, NewPlugin(c))
	if err != nil {
		return nil, err
	}
	return &Server{Server: s}, nil
}

type Plugin struct {
	c Core
}

func NewPlugin(c Core) *Plugin {
	return &Plugin{c: c}
}

func (p *Plugin) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return &PluginServer{core: p.c}, nil
}

func (_ *Plugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginClient{c: c}, nil
}

type PluginServer struct {
	core Core
}

var _ Core = (*PluginClient)(nil)

type PluginClient struct {
	c *rpc.Client
}

func (c *PluginClient) call(method string, args interface{}, reply interface{}) error {
	return c.c.Call("Plugin."+method, args, reply)
}

type StartParams struct {
	DataPath string
	Config   []byte
}

func (s *PluginServer) Start(p *StartParams, err *error) error {
	*err = s.core.Start(p.DataPath, p.Config)
	return nil
}
func (c *PluginClient) Start(dataPath string, config []byte) (err error) {
	err2 := c.call("Start", &StartParams{
		DataPath: dataPath,
		Config:   config,
	}, &err)
	if err2 != nil {
		return err
	}
	return nil
}

func (s *PluginServer) Close(_ interface{}, err *error) error {
	*err = s.core.Close()
	return nil
}
func (c *PluginClient) Close() (err error) {
	err2 := c.call("Close", new(interface{}), &err)
	if err2 != nil {
		return err
	}
	return nil
}

func (s *PluginServer) Protocols(_ interface{}, rsp *[]string) error {
	*rsp = s.core.Protocols()
	return nil
}
func (c *PluginClient) Protocols() (ps []string) {
	_ = c.call("Protocols", new(interface{}), &ps)
	return ps
}

func (s *PluginServer) Type(_ interface{}, t *string) error {
	*t = s.core.Type()
	return nil
}
func (c *PluginClient) Type() (t string) {
	_ = c.call("Type", new(interface{}), &t)
	return t
}
