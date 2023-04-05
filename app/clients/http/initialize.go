package http

import (
	"net"
	"net/netip"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

func GetInitialInfo(target netip.AddrPort, nodeIp net.IP, nodePort int, token string, csr []byte, serverId string) (*pb.RegisterResponse, error) {
	var request = &pb.RegisterRequest{
		Ip:       nodeIp.String(),
		Port:     int32(nodePort),
		Token:    token,
		ServerID: serverId,
		Csr:      csr,
	}
	var response = &pb.RegisterResponse{}

	client := netUtils.GetInsecureHttpClient()
	r := resty.NewWithClient(client)

	resp, err := r.R().
		SetHeader("Content-Type", "application/json").
		SetHeader(constant.RPCSourceMetadataKey, global.App.NodeInfo.ServerId).
		SetBody(request).
		SetResult(response).
		Post("https://" + target.String() + "/initialize")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New("failed to register node: " + resp.String())
	}

	response = resp.Result().(*pb.RegisterResponse)
	return response, nil
}
