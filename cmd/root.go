package cmd

import (
	"errors"
	"log"
	"net"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

var rootCmd = &cobra.Command{
	Use:   "openPanel",
	Short: "Distributed linux panel",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := netUtils.CheckPublicIp(*listenIp); err != nil {
			return err
		}
		if *listenPort < 1 || *listenPort > 65535 {
			return errors.New("Invalid port " + strconv.Itoa(*listenPort))
		}
		return nil
	},
}

var listenIp *net.IP
var listenPort *int

func init() {
	listenIp = rootCmd.PersistentFlags().IPP("ip", "i", constant.DefaultListenIp, "IP address to listen on")
	listenPort = rootCmd.PersistentFlags().IntP("port", "p", constant.DefaultListenPort, "Port to listen on")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("%#v", err)
	}
}
