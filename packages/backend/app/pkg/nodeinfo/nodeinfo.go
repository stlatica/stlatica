package nodeinfo

import "os"

const Protocol = "activitypub"

type NodeInfo interface {
	GetNodeInfoData() NodeInfoData
}

func NewNodeInfo() NodeInfo {
	return &nodeInfo{}
}

type nodeInfo struct {
	// ここには必要に応じてcontextとかloggerとかを持たせる
}

func (n *nodeInfo) GetNodeInfoData() NodeInfoData {
	return NodeInfoData{
		OpenRegistrations: os.Getenv("OPEN_REGISTRATIONS"), // 基本的には値を環境変数から読み込む形になるはず
		Protocols:         Protocol,                        // 固定値は定数とかで置いてもいいかも
	}
}

// NodeInfoData is the data of node info.
type NodeInfoData struct {
	OpenRegistrations string `json:"openRegistrations"`
	Protocols         string `json:"protocols"`
	// 以下nodeinfoの定義を追加する
}
