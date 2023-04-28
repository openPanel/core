package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

func NodesDbToPb(nodes []*shared.Node) []*pb.Node {
	pbNodes := make([]*pb.Node, len(nodes))
	for i, node := range nodes {
		pbNodes[i] = &pb.Node{
			Id:        node.ID,
			CreatedAt: timestamppb.New(node.CreatedAt),
			UpdatedAt: timestamppb.New(node.UpdatedAt),
			Name:      wrapperspb.String(node.Name),
			Ip:        node.IP,
			Port:      int64(node.Port),
			Comment:   wrapperspb.String(node.Comment),
		}
	}
	return pbNodes
}

func NodesPbToRouter(nodes []*pb.Node) []router.Node {
	routerNodes := make([]router.Node, len(nodes))
	for i, node := range nodes {
		routerNodes[i] = router.Node{
			Id: node.Id,
			AddrPort: netUtils.NewAddrPortWithString(
				node.Ip,
				int(node.Port),
			),
		}
	}
	return routerNodes
}

func NodesPbToDb(nodes []*pb.Node) []*shared.Node {
	dbNodes := make([]*shared.Node, len(nodes))
	for i, node := range nodes {
		dbNodes[i] = &shared.Node{
			ID:        node.Id,
			CreatedAt: node.CreatedAt.AsTime(),
			UpdatedAt: node.UpdatedAt.AsTime(),
			Name:      node.Name.Value,
			IP:        node.Ip,
			Port:      int(node.Port),
			Comment:   node.Comment.Value,
		}
	}
	return dbNodes
}
