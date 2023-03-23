package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var startLongHelp = `The start command is used to launch either a standalone application or the first node of a cluster.
The will issue the Certificate Authority (CA) certificate and key, which will be used to sign all other certificates in the cluster.`

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new panel instance",
	Long:  startLongHelp,
	Run:   startAction,
}

func startAction(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetIP("ip")
	port, _ := cmd.Flags().GetInt("port")

	if port < 1 || port > 65535 {
		panic("Invalid port" + strconv.Itoa(port))
	}

	_ = host
}

func init() {
	rootCmd.AddCommand(startCmd)

	attachServerFlags(startCmd)
}
