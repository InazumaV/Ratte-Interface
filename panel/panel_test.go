package panel

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

var p *PluginClient

func init() {
	// テスト用のサーバー実装をビルド
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
	p = cli
}

func TestPluginClient_AddRemote(t *testing.T) {
	// AddRemoteのテスト
	resp := p.AddRemote(&AddRemoteParams{
		Name:     "test-node",
		Baseurl:  "http://127.0.0.1:8080",
		NodeId:   1,
		NodeType: "model_context",
		Key:      "test-key",
		Timeout:  30,
	})
	if resp.Err != nil {
		t.Errorf("AddRemote error: %v", resp.Err)
	} else {
		t.Logf("AddRemote success: %d", resp.RemoteId)
	}
}

func TestPluginClient_GetNodeInfo(t *testing.T) {
	// GetNodeInfoのテスト
	nodeInfo := p.GetNodeInfo(1)
	log.Print(nodeInfo)
	if nodeInfo.Err != nil {
		t.Errorf("GetNodeInfo error: %v", nodeInfo.Err)
	} else {
		t.Logf("GetNodeInfo success: %v", nodeInfo.NodeInfo)
	}
}

func TestPluginClient_GetUserList(t *testing.T) {
	// GetUserListのテスト
	userList := p.GetUserList(1)
	if userList.Err != nil {
		t.Errorf("GetUserList error: %v", userList.Err)
	} else {
		t.Logf("GetUserList success: %d users", len(userList.Users))
	}
}

func TestPluginClient_ReportUserTraffic(t *testing.T) {
	// ReportUserTrafficのテスト
	err := p.ReportUserTraffic(&ReportUserTrafficParams{
		Id: 1,
		Users: []UserTrafficInfo{
			{
				Id:       1,
				Name:     "user1",
				Upload:   1024,
				Download: 2048,
			},
		},
	})
	if err != nil {
		t.Errorf("ReportUserTraffic error: %v", err)
	} else {
		t.Logf("ReportUserTraffic success")
	}
}

func TestPluginClient_DelRemote(t *testing.T) {
	// DelRemoteのテスト
	err := p.DelRemote(1)
	if err != nil {
		t.Errorf("DelRemote error: %v", err)
	} else {
		t.Logf("DelRemote success")
	}
}
