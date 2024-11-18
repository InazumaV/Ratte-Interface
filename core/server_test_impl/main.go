package main

import (
	"github.com/Yuzuki616/Ratte-Interface/core"
)

type impl struct {
}

func (i *impl) Start(dataPath string, config []byte) error {
	return nil
}
func (i *impl) Close() error {
	return nil
}
func (i *impl) AddNode(params *core.AddNodeParams) error {
	return nil
}
func (i *impl) DelNode(name string) error {
	return nil
}
func (i *impl) AddUsers(p *core.AddUsersParams) error {
	return nil
}
func (i *impl) GetUserTraffic(p *core.GetUserTrafficParams) *core.GetUserTrafficResponse {
	return nil
}
func (i *impl) ResetUserTraffic(p *core.ResetUserTrafficParams) error {
	return nil
}
func (i *impl) DelUsers(params *core.DelUsersParams) error {
	return nil
}
func (i *impl) Protocols() []string {
	return []string{
		"test",
	}
}
func (i *impl) Type() string {
	return "server"
}

func main() {
	s, err := core.NewServer(nil, new(impl))
	if err != nil {
		panic(err)
	}
	println(s.Run().Error())
}
