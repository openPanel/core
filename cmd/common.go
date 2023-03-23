package cmd

import (
	"net"

	"github.com/spf13/cobra"

	"github.com/openPanel/core/app/constant"
)

func attachServerFlags(cmd *cobra.Command) {
	cmd.Flags().IPP("ip", "i", net.ParseIP(constant.DefaultListenIp), "The IP address to bind to")
	cmd.Flags().IntP("port", "p", constant.DefaultListenPort, "The port to bind to")
	cmd.Flags().StringP("data", "d", constant.DefaultDataDir, "The data directory")
}
