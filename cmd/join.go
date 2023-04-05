package cmd

import (
	"net"

	"github.com/urfave/cli/v2"

	"github.com/openPanel/core/app/bootstrap"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

var joinCmd = &cli.Command{
	Name:  "join",
	Usage: "Join a existing panel cluster",
	Flags: []cli.Flag{
		&cli.GenericFlag{
			Name:  "ip",
			Value: NewIP(constant.DefaultListenIp),
			Usage: "IP address to listen on",
		},
		&cli.GenericFlag{
			Name:  "port",
			Value: NewPort(constant.DefaultListenPort),
			Usage: "Port to listen on",
		},
	},
	Description: "Enable a new node to join an existing distributed cluster. Address can belong to any node in the cluster.",
	ArgsUsage:   "<address> <token>",
	Action: func(context *cli.Context) error {
		if context.NArg() != 2 {
			return cli.ShowCommandHelp(context, "join")
		}
		address := context.Args().Get(0)
		token := context.Args().Get(1)

		ip, port := netUtils.AssertPublicAddress(address)

		listenIP := context.Generic("ip").(*IP)
		listenPort := context.Generic("port").(*Port)

		bootstrap.Join(net.IP(*listenIP), int(*listenPort), ip, port, token)

		return nil
	},
}
