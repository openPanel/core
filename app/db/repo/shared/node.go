package shared

import (
	"context"

	"github.com/openPanel/core/app/db/db"
	"github.com/openPanel/core/app/generated/db/shared"
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
