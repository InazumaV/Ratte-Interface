package panel

import (
	"github.com/Yuzuki616/Ratte-Interface/params"
)

type NodeInfo params.NodeInfo

type GetNodeInfoRsp struct {
	Hash     string
	NodeInfo NodeInfo
	Err      error
}

func (s *PluginServer) GetNodeInfo(id int, r *GetNodeInfoRsp) error {
	*r = *s.p.GetNodeInfo(id)
	return nil
}

func (c *PluginClient) GetNodeInfo(id int) (r *GetNodeInfoRsp) {
	r = &GetNodeInfoRsp{}
	err := c.call("GetNodeInfo", id, r)
	if err != nil {
		r.Err = err
	}
	return nil
}
