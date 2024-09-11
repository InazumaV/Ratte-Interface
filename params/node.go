package params

import (
	"encoding/json"
	"fmt"
)

type CommonNodeInfo struct {
	Type        string
	VMess       *VMessNode
	VLess       *VLessNode
	Shadowsocks *ShadowsocksNode
	Trojan      *TrojanNode
	Hysteria    *HysteriaNode
	Other       *OtherNode
}

func (n *CommonNodeInfo) String() string {
	switch n.Type {
	case "VMess":
		return fmt.Sprintf("VMess: %v", n.VMess)
	case "Shadowsocks":
		return fmt.Sprintf("Shadowsocks: %v", n.Shadowsocks)
	case "Trojan":
		return fmt.Sprintf("Trojan: %v", n.Trojan)
	case "Hysteria":
		return fmt.Sprintf("Hysteria: %v", n.Hysteria)
	case "Other":
		return fmt.Sprintf("Other: %v", n.Other)
	default:
		return fmt.Sprintf("%v", *n)
	}
}

type CommonNode struct {
	Name          string
	Host          string
	Port          string
	ProxyProtocol bool
	TCPFastOpen   bool
	EnableDNS     bool
	Rules         []string
	Limit         LimitOptions
	ExpandParams
}

type HysteriaNode struct {
	CommonNode
	UpMbps   int
	DownMbps int
	Obfs     string
}

// VMessNode is vmess node info
type VMessNode struct {
	CommonNode
	TlsType         int
	Network         string
	ServerName      string
	TlsSettings     TlsSettings
	NetworkSettings json.RawMessage
}

// VLessNode is vless node info
type VLessNode struct {
	CommonNode
	TlsType         int
	Flow            string
	Network         string
	ServerName      string
	RealityConfig   RealityConfig
	TlsSettings     TlsSettings
	NetworkSettings json.RawMessage
}

type TlsSettings struct {
	ServerName string
	ServerPort string
	ShortId    string
	PrivateKey string
}

type RealityConfig struct {
	Xver         uint64 `json:"Xver"`
	MinClientVer string `json:"MinClientVer"`
	MaxClientVer string `json:"MaxClientVer"`
	MaxTimeDiff  string `json:"MaxTimeDiff"`
}

type ShadowsocksNode struct {
	CommonNode
	Cipher    string `json:"cipher"`
	ServerKey string `json:"server_key"`
}

type TrojanNode CommonNode

type OtherNode struct {
	Name    string
	TlsType int
	CommonNode
	ExpandParams
}

type LimitOptions struct {
	SpeedLimit int `json:"SpeedLimit"`
	IPLimit    int `json:"DeviceLimit"`
	ConnLimit  int `json:"ConnLimit"`
}
