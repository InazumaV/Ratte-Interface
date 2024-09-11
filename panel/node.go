package panel

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/Yuzuki616/Ratte-Interface/params"
)

type NodeInfo params.NodeInfo

type GetNodeInfoRsp struct {
	Hash     string
	NodeInfo NodeInfo
	Err      error
}

func (g *GetNodeInfoRsp) GetHash() string {
	if len(g.Hash) > 0 {
		return g.Hash
	}
	s := sha256.Sum256([]byte(fmt.Sprintf("%x", fmt.Sprintf("%v", g.NodeInfo))))
	return hex.EncodeToString(s[:])
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
