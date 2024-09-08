package panel

import (
	"github.com/Yuzuki616/Ratte-Interface/params"
)

type UserInfo params.UserInfo

type GetUserListRsp struct {
	Users []UserInfo
	Err   error
}

func (s *PluginServer) GetUserList(id int, r *GetUserListRsp) error {
	*r = *s.p.GetUserList(id)
	return nil
}
func (c *PluginClient) GetUserList(id int) (r *GetUserListRsp) {
	r = &GetUserListRsp{}
	err := c.call("GetUserList", id, r)
	if err != nil {
		r.Err = err
	}
	return
}

type ReportUserTrafficParams struct {
	Id    int
	Users []UserTrafficInfo
}
type UserTrafficInfo struct {
	Id       int
	Name     string
	Upload   int64
	Download int64
}

func (s *PluginServer) ReportUserTraffic(p *ReportUserTrafficParams, err *error) error {
	*err = s.p.ReportUserTraffic(p)
	return nil
}
func (c *PluginClient) ReportUserTraffic(p *ReportUserTrafficParams) (err error) {
	err2 := c.call("ReportUserTraffic", p, &err)
	if err2 != nil {
		return err2
	}
	return
}
