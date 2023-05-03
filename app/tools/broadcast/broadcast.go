package broadcast

import (
	"encoding/json"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/convert"
)

// Things about serialization and deserialization of broadcast message
type linkStatesBroadcastMessage struct {
	LinkStates []*pb.LinkState `json:"link_states"`
	AddNodes   *[]router.Node  `json:"add_nodes,omitempty"`
	UpdateNode *[]router.Node  `json:"update_node,omitempty"`
	DelNodes   *[]router.Node  `json:"del_nodes,omitempty"`
}

func GetRouterPayload(lst router.LinkStates, add *[]router.Node, update *[]router.Node, del *[]router.Node) ([]byte, error) {
	bm := linkStatesBroadcastMessage{
		LinkStates: convert.LinkStatesRouterToPb(lst),
	}
	if add != nil {
		bm.AddNodes = add
	}
	if update != nil {
		bm.UpdateNode = update
	}
	if del != nil {
		bm.DelNodes = del
	}

	return json.Marshal(bm)
}

func LoadRouterPayload(payload []byte) error {
	bm := new(linkStatesBroadcastMessage)
	err := json.Unmarshal(payload, bm)
	if err != nil {
		return err
	}

	router.EditNodes(func(ns []router.Node) []router.Node {
		if bm.AddNodes != nil {
			ns = append(ns, *bm.AddNodes...)
		}
		if bm.UpdateNode != nil {
			for _, n := range *bm.UpdateNode {
				for i, nn := range ns {
					if nn.Id == n.Id {
						ns[i] = n
						break
					}
				}
			}
		}
		if bm.DelNodes != nil {
			for _, n := range *bm.DelNodes {
				for i, nn := range ns {
					if nn.Id == n.Id {
						ns = append(ns[:i], ns[i+1:]...)
						break
					}
				}
			}
		}

		return ns
	})

	ls := convert.LinkStatesPbToRouter(bm.LinkStates)
	router.UpdateLinkStates(ls)

	return nil
}
