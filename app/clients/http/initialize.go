package http

import (
	"fmt"
	"net"
	"net/netip"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

func getAuthedRequest(token string, serverId string) *resty.Request {
	client := netUtils.GetInsecureHttpClient()
	r := resty.NewWithClient(client).R()
	r.SetHeader(constant.HttpContentTypeHeader, constant.ContentTypeJson)
	r.SetHeader(constant.HttpAuthorizationTokenHeader, token)
	if serverId != "" {
		r.SetHeader(constant.RPCSourceMetadataKey, serverId)
	}
	return r
}

func GetClusterInfo(target netip.AddrPort, token string) (*pb.GetClusterInfoResponse, error) {
	var response = &pb.GetClusterInfoResponse{}

	resp, err := getAuthedRequest(token, "").
		SetResult(response).
		Get(fmt.Sprintf("https://%s/initialize", target.String()))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New("failed to get cluster info: " + resp.String())
	}

	response = resp.Result().(*pb.GetClusterInfoResponse)
	return response, nil
}

func RegisterNewNode(
	target netip.AddrPort,
	nodeIp net.IP,
	nodePort int,
	token string,
	csr []byte,
	serverId string,
	linkStates []*pb.LinkState) (*pb.RegisterResponse, error) {
	request := &pb.RegisterRequest{
		Ip:         nodeIp.String(),
		Port:       int32(nodePort),
		ServerID:   serverId,
		Csr:        csr,
		LinkStates: linkStates,
	}
	var response = &pb.RegisterResponse{}

	resp, err := getAuthedRequest(token, serverId).
		SetBody(request).
		SetResult(response).
		Post(fmt.Sprintf("https://%s/initialize", target.String()))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New("failed to register node: " + resp.String())
	}

	response = resp.Result().(*pb.RegisterResponse)
	return response, nil
}
