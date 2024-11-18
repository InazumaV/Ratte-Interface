package core

import "github.com/Yuzuki616/Ratte-Interface/params"

type NodeInfo struct {
	params.CommonNodeInfo
	TlsOptions
}
type TlsOptions struct {
	CertPath string
	KeyPath  string
}

type AddNodeParams struct {
	NodeInfo
	ExpandOptions []byte
}

func (s *PluginImplServer) AddNode(params *AddNodeParams, err *error) error {
	*err = s.core.AddNode(params)
	return nil
}
func (c *PluginImplClient) AddNode(params *AddNodeParams) (err error) {
	err2 := c.call("AddNode", params, &err)
	if err2 != nil {
		return err2
	}
	return
}

func (s *PluginImplServer) DelNode(name string, err *error) error {
	*err = s.core.DelNode(name)
	return nil
}

func (c *PluginImplClient) DelNode(name string) (err error) {
	err2 := c.call("DelNode", name, &err)
	if err2 != nil {
		return err2
	}
	return
}
