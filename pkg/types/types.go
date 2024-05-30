package types

type VswitchPostBody struct {
	Bridge   string         `json:"bridge"`
	Topology []MeshTopology `json:"topology"`
}

type MeshTopology struct {
	NodeIP string `json:"nodeIP"`
	VNI    int32  `json:"vni"`
}