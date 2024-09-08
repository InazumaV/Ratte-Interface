package panel

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

const PluginName = "ratte-panel"

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "is-ratte-panel",
	MagicCookieValue: "true",
}

type Panel interface {
	AddRemote(params *AddRemoteParams) *AddRemoteRsp
	DelRemote(id int) error
	GetNodeInfo(id int) *GetNodeInfoRsp
	GetUserList(id int) *GetUserListRsp
	ReportUserTraffic(p *ReportUserTrafficParams) error
}

type Plugin struct {
	p Panel
}

func NewPlugin(impl Panel) *Plugin {
	return &Plugin{
		p: impl,
	}
}

func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &PluginServer{p: p.p}, nil
}

func (_ *Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginClient{c: c}, nil
}

type PluginServer struct {
	p Panel
}

var _ Panel = (*PluginClient)(nil)

type PluginClient struct {
	c *rpc.Client
}

func (c *PluginClient) call(method string, args interface{}, reply interface{}) error {
	return c.c.Call("Plugin."+method, args, reply)
}
