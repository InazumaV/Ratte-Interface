package panel

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/InazumaV/Ratte-Interface/common/errors"
	"github.com/InazumaV/Ratte-Interface/params"
)

type UserInfo struct {
	HashOrKey string
	params.UserInfo
}

func (u *UserInfo) GetHashOrKey() string {
	if len(u.HashOrKey) != 0 {
		return u.HashOrKey
	}
	s := sha256.Sum256([]byte(fmt.Sprintf("%x", fmt.Sprintf("%v", u))))
	return hex.EncodeToString(s[:])
}

type GetUserListRsp struct {
	Hash  string
	Users []UserInfo
	Err   error
}

func (g *GetUserListRsp) GetHash() string {
	if len(g.Hash) > 0 {
		return g.Hash
	}
	s := sha256.Sum256([]byte(fmt.Sprintf("%x", fmt.Sprintf("%v", g.Users))))
	return hex.EncodeToString(s[:])
}

func (s *PluginImplServer) GetUserList(id int, r *GetUserListRsp) error {
	*r = *s.p.GetUserList(id)
	if r.Err != nil {
		r.Err = errors.NewStringFromErr(r.Err)
	}
	return nil
}
func (c *PluginImplClient) GetUserList(id int) (r *GetUserListRsp) {
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

func (s *PluginImplServer) ReportUserTraffic(p *ReportUserTrafficParams, err *error) error {
	*err = errors.NewStringFromErr(s.p.ReportUserTraffic(p))
	return nil
}
func (c *PluginImplClient) ReportUserTraffic(p *ReportUserTrafficParams) (err error) {
	err2 := c.call("ReportUserTraffic", p, &err)
	if err2 != nil {
		return err2
	}
	return
}
