package main

import (
	"github.com/InazumaV/Ratte-Interface/panel"
)

type impl struct {
}

func (i *impl) AddRemote(params *panel.AddRemoteParams) *panel.AddRemoteRsp {
	return &panel.AddRemoteRsp{
		RemoteId: 1,
		Err:      nil,
	}
}

func (i *impl) DelRemote(id int) error {
	return nil
}

func (i *impl) GetNodeInfo(id int) *panel.GetNodeInfoRsp {
	// モデルコンテキストプロトコルのノード情報をサンプルとして返す
	return &panel.GetNodeInfoRsp{}
}

func (i *impl) GetUserList(id int) *panel.GetUserListRsp {
	// サンプルユーザーのリストを返す
	return &panel.GetUserListRsp{}
}

func (i *impl) ReportUserTraffic(p *panel.ReportUserTrafficParams) error {
	return nil
}

func main() {
	s, err := panel.NewServer(nil, new(impl))
	if err != nil {
		panic(err)
	}
	println(s.Run().Error())
}
