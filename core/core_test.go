package core

import (
	"os/exec"
	"testing"
)

var c *PluginClient

func init() {
	exec.Command("go", "build", "./server_test_impl/main.go", "-o", "./server_test_impl/server_test_impl")
	cli, err := NewClient(nil, exec.Command("./server_test_impl/server_test_impl"))
	if err != nil {
		panic(err)
	}
	c = cli
}

func TestPluginClient_Type(t *testing.T) {
	t.Log(c.Type())
}

func TestPluginClient_DelUsers(t *testing.T) {
	t.Log(c.DelUsers(new(DelUsersParams)))
}
