package cmd

import (
	"net"

	"github.com/urfave/cli/v2"

	"github.com/openPanel/core/app/bootstrap"
	"github.com/openPanel/core/app/constant"
)

var startLongHelp = `The start command is used to launch either a standalone application or the first node of a cluster.
The will issue the Certificate Authority (CA) certificate and key, which will be used to sign all other certificates in the cluster.`

//
//var startCmd = &cobra.Command{
//	Use:     "start",
//	Short:   "Start a new panel instance",
//	Long:    startLongHelp,
//	PreRunE: preRunCheck,
//	Run:     startAction,
//}
//
//func startAction(cmd *cobra.Command, _ []string) {
//	bootstrap.Start(*listenIp, *listenPort)
//}
//
//func init() {
//	rootCmd.AddCommand(startCmd)
//
//	attachAddress(startCmd)
//}

var startCmd = &cli.Command{
	Name:  "start",
	Usage: "Start a new panel instance",
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
	Description: startLongHelp,
	Action: func(context *cli.Context) error {
		listenIP := context.Generic("ip").(*IP)
		listenPort := context.Generic("port").(*Port)

		bootstrap.Start(net.IP(*listenIP), int(*listenPort))

		return nil
	},
}
