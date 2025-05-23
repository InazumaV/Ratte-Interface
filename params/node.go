package params

import (
	"fmt"
)

type NodeInfo struct {
	ExpandParams
	Type           string
	Name           string
	Port           int
	ProxyProtocol  bool
	TCPFastOpen    bool
	EnableDNS      bool
	Security       string
	SecurityConfig *SecurityConfig
	Limit          LimitOptions
	Rules          []string

	VMess       *VMess
	VLess       *VLess
	Shadowsocks *Shadowsocks
	Trojan      *Trojan
	Hysteria    *Hysteria
	Other       *Other
}

func (n *NodeInfo) String() string {
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

type Hysteria struct {
	ExpandParams
	UpMbps   int
	DownMbps int
	Obfs     string
}

// VMess is vmess node info
type VMess struct {
	ExpandParams
	NetworkSettings
	Network string
}

type NetworkSettings struct {
	Ws   WsSettings
	Grpc GrpcSettings
}

type WsSettings struct {
	Headers         map[string]string
	Host            string
	Path            string
	HeartbeatPeriod string
}

type GrpcSettings struct {
	Authority           string
	ServiceName         string
	MultiMode           bool
	UserAgent           string
	IdleTimeout         int
	HealthCheckTimeout  int
	PermitWithoutStream bool
	InitialWindowsSize  int
}

// VLess is vless node info
type VLess struct {
	VMess
	Flow string
}

type TlsSettings struct {
	ServerName    string
	AllowInsecure bool   `json:"allow_insecure"`
	Fingerprint   string `json:"fingerprint"`
}

type SecurityConfig struct {
	TlsSettings   TlsSettings
	RealityConfig RealityConfig
}

type RealityConfig struct {
	ServerName   string
	ServerPort   int
	ShortId      string
	PrivateKey   string
	Xver         uint64
	MinClientVer string
	MaxClientVer string
	MaxTimeDiff  string
}

type Shadowsocks struct {
	ExpandParams
	Cipher    string
	ServerKey string
}

type Trojan struct {
	ExpandParams
	Host string
	Path string
}

type Other ExpandParams

type LimitOptions struct {
	SpeedLimit uint64 `json:"SpeedLimit"`
	IPLimit    int    `json:"DeviceLimit"`
	ConnLimit  int    `json:"ConnLimit"`
}
