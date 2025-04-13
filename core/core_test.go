package core

import (
	"os"
	"os/exec"
	"testing"
)

var c *PluginClient

func init() {
	e := exec.Command("go", "build", "-C", "./server_test_impl/", "-o", "server_test_impl")
	e.Stdout = os.Stdout
	e.Stderr = os.Stderr
	if err := e.Run(); err != nil {
		panic(err)
	}
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
