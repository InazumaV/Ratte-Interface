package panel

import "testing"

func TestPluginImplClient_DelRemote(t *testing.T) {
	err := p.DelRemote(0)
	if err != nil {
		t.Fatal(err)
	}
}
