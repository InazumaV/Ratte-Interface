package core

import (
	params2 "github.com/Yuzuki616/Ratte-Interface/params"
)

type AddUsersParams struct {
	NodeName string
	Users    []params2.UserInfo
	params2.ExpandParams
}

func (s *PluginServer) AddUsers(params *AddUsersParams, err *error) error {
	*err = s.core.AddUsers(params)
	return nil
}
func (c *PluginClient) AddUsers(p *AddUsersParams) (err error) {
	err2 := c.call("AddUsers", p, &err)
	if err2 != nil {
		return err2
	}
	return
}

type GetUserTrafficParams struct {
	NodeName string
	Username string
}
type GetUserTrafficResponse struct {
	up   int64
	down int64
	Err  error
}

func (s *PluginServer) GetUserTraffic(params *GetUserTrafficParams, rsp *GetUserTrafficResponse) error {
	*rsp = *s.core.GetUserTraffic(params)
	return nil
}
func (c *PluginClient) GetUserTraffic(p *GetUserTrafficParams) (rsp *GetUserTrafficResponse) {
	rsp = &GetUserTrafficResponse{}
	err := c.call("GetUserTraffic", p, rsp)
	if err != nil {
		rsp = &GetUserTrafficResponse{Err: err}
	}
	return
}

type ResetUserTrafficParams struct {
	NodeName string
	Username string
}

func (s *PluginServer) ResetUserTraffic(p *ResetUserTrafficParams, err *error) error {
	*err = s.core.ResetUserTraffic(p)
	return nil
}
func (c *PluginClient) ResetUserTraffic(p *ResetUserTrafficParams) (err error) {
	return c.call("ResetUserTraffic", p, &err)
}

type DelUsersParams struct {
	NodeName string
	Users    []string
}

func (s *PluginServer) DelUsers(params *DelUsersParams, err *error) error {
	*err = s.core.DelUsers(params)
	return nil
}
func (c *PluginClient) DelUsers(params *DelUsersParams) (err error) {
	return c.call("DelUsers", params, &err)
}
