package panel

import (
	"github.com/InazumaV/Ratte-Interface/params"
)

type AddRemoteParams struct {
	Baseurl  string
	NodeId   int
	NodeType string
	Key      string
	Timeout  int
	params.ExpandParams
}

type AddRemoteRsp struct {
	RemoteId int
	Err      error
}

func (s *PluginImplServer) AddRemote(params *AddRemoteParams, r *AddRemoteRsp) error {
	*r = *s.p.AddRemote(params)
	return nil
}
func (c *PluginImplClient) AddRemote(params *AddRemoteParams) (r *AddRemoteRsp) {
	r = &AddRemoteRsp{}
	err := c.call("AddRemote", params, r)
	if err != nil {
		r.Err = err
	}
	return
}

func (s *PluginImplServer) DelRemote(id int, _ any) {
	_ = s.p.DelRemote(id)
}

func (c *PluginImplClient) DelRemote(id int) error {
	err := c.call("DelRemote", id, new(any))
	if err != nil {
		return err
	}
	return nil
}
