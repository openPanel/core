package shared

import (
	"context"

	"github.com/openPanel/core/app/db/db"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/global"
)

type nodeRepo struct{}

var NodeRepo = new(nodeRepo)

func (r *nodeRepo) AddNode(ctx context.Context, serverId string, ip string, port int) error {
	return db.GetSharedDb().Node.
		Create().
		SetID(serverId).
		SetIP(ip).
		SetPort(port).
		Exec(ctx)
}

func (r *nodeRepo) GetAll(ctx context.Context) ([]*shared.Node, error) {
	return db.GetSharedDb().Node.Query().All(ctx)
}

func (r *nodeRepo) GetBroadcastNodes(ctx context.Context) ([]*shared.Node, error) {
	nodes, err := db.GetSharedDb().Node.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	broadcastNodes := make([]*shared.Node, 0, len(nodes)-1)
	for _, node := range nodes {
		if node.ID != global.App.NodeInfo.ServerId {
			broadcastNodes = append(broadcastNodes, node)
		}
	}
	return broadcastNodes, nil
}
