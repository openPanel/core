package router

import (
	"encoding/json"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/tools/convert"
)

// Things about serialization and deserialization of broadcast message
type broadcastMessage struct {
	LinkStates []*pb.LinkState `json:"link_states"`
	Nodes      []Node          `json:"nodes"`
}

func GetBroadcastPayload() ([]byte, error) {
	bm := broadcastMessage{
		LinkStates: convert.LinkStatesRouterToPb(GetLinkStates()),
		Nodes:      flattenNodes(),
	}
	return json.Marshal(bm)
}

func LoadBroadcastPayload(payload []byte) error {
	bm := new(broadcastMessage)
	err := json.Unmarshal(payload, bm)
	if err != nil {
		return err
	}

	ls := convert.LinkStatesPbToRouter(bm.LinkStates)
	ns := make([]Node, 0, len(bm.Nodes))
	for _, n := range bm.Nodes {
		ns = append(ns, Node{
			Id:       n.Id,
			AddrPort: n.AddrPort,
		})
	}

	SetNodes(ns)
	MergeLinkStates(ls)

	return nil
}
