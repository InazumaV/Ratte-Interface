package core

import (
	params2 "github.com/Yuzuki616/Ratte-Interface/params"
)

type UserInfo params2.UserInfo
type AddUsersParams struct {
	NodeName string
	Users    []UserInfo
	params2.ExpandParams
}

func (s *PluginImplServer) AddUsers(params *AddUsersParams, err *error) error {
	*err = s.core.AddUsers(params)
	return nil
}
func (c *PluginImplClient) AddUsers(p *AddUsersParams) (err error) {
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
	Up   int64
	Down int64
	Err  error
}

func (s *PluginImplServer) GetUserTraffic(params *GetUserTrafficParams, rsp *GetUserTrafficResponse) error {
	*rsp = *s.core.GetUserTraffic(params)
	return nil
}
func (c *PluginImplClient) GetUserTraffic(p *GetUserTrafficParams) (rsp *GetUserTrafficResponse) {
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

func (s *PluginImplServer) ResetUserTraffic(p *ResetUserTrafficParams, err *error) error {
	*err = s.core.ResetUserTraffic(p)
	return nil
}
func (c *PluginImplClient) ResetUserTraffic(p *ResetUserTrafficParams) (err error) {
	return c.call("ResetUserTraffic", p, &err)
}

type DelUsersParams struct {
	NodeName string
	Users    []string
}

func (s *PluginImplServer) DelUsers(params *DelUsersParams, err *error) error {
	*err = s.core.DelUsers(params)
	return nil
}
func (c *PluginImplClient) DelUsers(params *DelUsersParams) (err error) {
	return c.call("DelUsers", params, &err)
}
