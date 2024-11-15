package main

import (
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/plugin"
)

type impl struct {
	core.Core
}

func (i *impl) Type() string {
	return "server"
}

func main() {
	s, err := plugin.NewServer(nil, core.Core(new(impl)))
	if err != nil {
		panic(err)
	}
	println(s.Run().Error())
}
