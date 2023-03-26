package local

import (
	"context"

	"github.com/openPanel/core/app/db/db"
	"github.com/openPanel/core/app/generated/db/local"
)

type nodeRepo struct{}

var NodeRepo = new(nodeRepo)

// Clear Just clear the node table
func (r *nodeRepo) Clear(ctx context.Context) error {
	_, err := db.GetLocalDb().Node.Delete().Exec(ctx)
	return err
}

func (r *nodeRepo) Init(ctx context.Context, nodes []*local.Node) error {
	err := r.Clear(ctx)
	if err != nil {
		return err
	}
	creates := make([]*local.NodeCreate, 0)
	for _, node := range nodes {
		creates = append(creates, db.GetLocalDb().Node.Create().
			SetID(node.ID).
			SetName(node.Name).
			SetIP(node.IP).
			SetPort(node.Port).
			SetComment(node.Comment))
	}
	return db.GetLocalDb().Node.CreateBulk(creates...).Exec(ctx)
}

func (r *nodeRepo) GetAll(ctx context.Context) ([]*local.Node, error) {
	return db.GetLocalDb().Node.Query().All(ctx)
}

func (r *nodeRepo) SyncSharedNodes(ctx context.Context) error {
	nodes, err := db.GetSharedDb().Node.Query().All(ctx)
	if err != nil {
		return err
	}

	creates := make([]*local.NodeCreate, 0)
	for _, node := range nodes {
		creates = append(creates, db.GetLocalDb().Node.Create().
			SetID(node.ID).
			SetName(node.Name).
			SetIP(node.IP).
			SetPort(node.Port).
			SetComment(node.Comment))
	}
	return db.GetLocalDb().Node.CreateBulk(creates...).Exec(ctx)
}
