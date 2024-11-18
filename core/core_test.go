package core

import (
	"os/exec"
	"testing"
)

var c *Client

func init() {
	cli, err := NewClient(nil, exec.Command("./server_test_impl/server_test_impl"))
	if err != nil {
		panic(err)
	}
	c = cli
}

func TestPluginClient_Type(t *testing.T) {
	t.Log(c.Type())
}
