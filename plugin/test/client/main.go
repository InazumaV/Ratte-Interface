package main

import (
	"fmt"
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/plugin"
	"os/exec"
)

func main() {
	c, err := plugin.NewClient(plugin.TypeCore, &plugin.Config{
		Cmd: exec.Command("../server/server"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.Caller().(core.Core).Type())

}
