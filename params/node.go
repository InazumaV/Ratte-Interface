package params

import "encoding/json"

type NodeInfo struct {
	Type        string
	VAllss      *VAllssNode
	Shadowsocks *ShadowsocksNode
	Trojan      *TrojanNode
	Hysteria    *HysteriaNode
	Other       *OtherNode
	ExpandParams
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
}

type HysteriaNode struct {
	CommonNode
	UpMbps   int
	DownMbps int
	Obfs     string
}

// VAllssNode is vmess and vless node info
type VAllssNode struct {
	CommonNode
	Tls                 int             `json:"tls"`
	TlsSettings         TlsSettings     `json:"tls_settings"`
	TlsSettingsBack     *TlsSettings    `json:"tlsSettings"`
	Network             string          `json:"network"`
	NetworkSettings     json.RawMessage `json:"network_settings"`
	NetworkSettingsBack json.RawMessage `json:"networkSettings"`
	ServerName          string          `json:"server_name"`

	// vless only
	Flow          string        `json:"flow"`
	RealityConfig RealityConfig `json:"-"`
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
	CommonNode
	ExpandParams
}

type LimitOptions struct {
	SpeedLimit int `json:"SpeedLimit"`
	IPLimit    int `json:"DeviceLimit"`
	ConnLimit  int `json:"ConnLimit"`
}
