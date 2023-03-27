package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/openPanel/core/app/bootstrap"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

var startLongHelp = `The start command is used to launch either a standalone application or the first node of a cluster.
The will issue the Certificate Authority (CA) certificate and key, which will be used to sign all other certificates in the cluster.`

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new panel instance",
	Long:  startLongHelp,
	Run:   startAction,
}

func startAction(cmd *cobra.Command, _ []string) {
	if !listenIp.IsUnspecified() {
		err := netUtils.CheckPublicIp(*listenIp)
		if err != nil {
			log.Fatal(err)
		}
	}

	bootstrap.Start(*listenIp, *listenPort)
}

func init() {
	rootCmd.AddCommand(startCmd)
}
