package cmd

import (
	"github.com/spf13/cobra"

	"github.com/openPanel/core/app/bootstrap"
)

var startLongHelp = `The start command is used to launch either a standalone application or the first node of a cluster.
The will issue the Certificate Authority (CA) certificate and key, which will be used to sign all other certificates in the cluster.`

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new panel instance",
	Long:  startLongHelp,
	Run:   startAction,
}

func startAction(_ *cobra.Command, _ []string) {
	bootstrap.Start()
}

func init() {
	rootCmd.AddCommand(startCmd)
}
