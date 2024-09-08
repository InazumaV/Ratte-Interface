package panel

import (
	"github.com/Yuzuki616/Ratte-Interface/params"
)

type AddRemoteParams struct {
	Baseurl  string
	NodeId   int
	NodeType string
	Timeout  int
	params.ExpandParams
}

type AddRemoteRsp struct {
	RemoteId int
	Err      error
}

func (s *PluginServer) AddRemote(params *AddRemoteParams, r *AddRemoteRsp) error {
	*r = *s.p.AddRemote(params)
	return nil
}
func (c *PluginClient) AddRemote(params *AddRemoteParams) (r *AddRemoteRsp) {
	r = &AddRemoteRsp{}
	err := c.call("AddRemote", params, r)
	if err != nil {
		r.Err = err
	}
	return
}
