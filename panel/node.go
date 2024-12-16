package panel

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/InazumaV/Ratte-Interface/params"
)

const (
	NoTls = 0
	Tls   = 1
)

type NodeInfo params.NodeInfo

func (i *NodeInfo) TlsType() int {
	switch i.Type {
	case "vmess":
		if i.VMess.TlsType == Tls {
			return Tls
		}
	case "vless":
		if i.VLess.TlsType == Tls {
			return Tls
		}
		return NoTls
	case "trojan":
		return Tls
	case "hysteria":
		return Tls
	case i.Other.Name:
		return i.Other.TlsType
	default:
		return NoTls
	}
	return NoTls
}

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

func (s *PluginImplServer) GetNodeInfo(id int, r *GetNodeInfoRsp) error {
	*r = *s.p.GetNodeInfo(id)
	return nil
}

func (c *PluginImplClient) GetNodeInfo(id int) (r *GetNodeInfoRsp) {
	r = &GetNodeInfoRsp{}
	err := c.call("GetNodeInfo", id, r)
	if err != nil {
		r.Err = err
	}
	return nil
}
