package nodeinfo

import (
	"os"
)

// NodeInfoData
const PROTOCOL = "activitypub"
const NODEINFO_VERSION = "2.1"

// Software
const NAME = "stlatica"
const REPOSITORY = "https://github.com/stlatica/stlatic"
const SOFTWARE_VERSION = "0.1.0"

// Users
const TEST_USER_COUNT = 1

type NodeInfo interface {
	GetNodeInfoData() NodeInfoData
}

func NewNodeInfo() NodeInfo {
	return &nodeInfo{}
}

type nodeInfo struct {
	// ここには必要に応じてcontextとかloggerとかを持たせる
	NodeInfoData
}

func (n *nodeInfo) GetNodeInfoData() NodeInfoData {
	return NodeInfoData{
		OpenRegistrations: os.Getenv("OPEN_REGISTRATIONS"), // 基本的には値を環境変数から読み込む形になるはず
		Protocols:         PROTOCOL,                        // 固定値は定数とかで置いてもいいかも
		Software: Software{
			Name:       NAME,
			Repository: REPOSITORY,
			Version:    NODEINFO_VERSION,
		},
		Usage: Usage{
			Users: Users{
				//TODO: ユーザ数のカウントに置き換え　https://github.com/stlatica/stlatica/issues/318
				Total: TEST_USER_COUNT,
			},
		},
		Services: Services{
			Inbound:  []string{},
			Outbound: []string{},
		},
		Metadata: getMetadata(),
		Version:  SOFTWARE_VERSION,
	}
}

// NodeInfoData is the data of node info.
type NodeInfoData struct {
	OpenRegistrations string   `json:"openRegistrations"`
	Protocols         string   `json:"protocols"`
	Software          Software `json:"software"`
	Usage             Usage    `json:"usage"`
	Services          Services `json:"services"`
	Metadata          []string `json:"metadata"`
	Version           string   `json:"version"`
}

func getMetadata() []string {
	return []string{}
}

type Software struct {
	Name       string
	Repository string
	Version    string
}

type Usage struct {
	Users Users
}

type Users struct {
	Total int
}

type Services struct {
	Inbound  []string
	Outbound []string
}
