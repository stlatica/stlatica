package nodeinfo

import (
	"os"
)

// Protocol is member of NodeInfoData
const Protocol = "activitypub"

// NodeInfoVersion is member of Software
const NodeInfoVersion = "2.1"

// Name and following is member of Software
const Name = "stlatica"

// Repository is member of Software
const Repository = "https://github.com/stlatica/stlatic"

// SoftwareVersion is member of Software
const SoftwareVersion = "0.1.0"

// NodeInfo is interface for NodeInfoData
type NodeInfo interface {
	GetNodeInfoData(userCount int) NodeInfoData
}

// NewNodeInfo is constructor for NodeInfo
func NewNodeInfo() NodeInfo {
	return &nodeInfo{}
}

// nodeInfo is struct for NodeInfo
type nodeInfo struct {
	// ここには必要に応じてcontextとかloggerとかを持たせる
	NodeInfoData
}

func (n *nodeInfo) GetNodeInfoData(userCount int) NodeInfoData {
	return NodeInfoData{
		OpenRegistrations: os.Getenv("OPEN_REGISTRATIONS"), // 基本的には値を環境変数から読み込む形になるはず
		Protocols:         Protocol,                        // 固定値は定数とかで置いてもいいかも
		Software: software{
			Name:       Name,
			Repository: Repository,
			Version:    NodeInfoVersion,
		},
		Usage: usage{
			Users: users{
				Total: userCount,
			},
		},
		Services: services{
			Inbound:  []string{},
			Outbound: []string{},
		},
		Metadata: getMetadata(),
		Version:  SoftwareVersion,
	}
}

// NodeInfoData is the data of node info.
type NodeInfoData struct {
	OpenRegistrations string   `json:"openRegistrations"`
	Protocols         string   `json:"protocols"`
	Software          software `json:"software"`
	Usage             usage    `json:"usage"`
	Services          services `json:"services"`
	Metadata          []string `json:"metadata"`
	Version           string   `json:"version"`
}

func getMetadata() []string {
	return []string{}
}

type software struct {
	Name       string
	Repository string
	Version    string
}

type usage struct {
	Users users
}

type users struct {
	Total int
}

type services struct {
	Inbound  []string
	Outbound []string
}
