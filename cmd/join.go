package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/openPanel/core/app/tools/utils"
)

var joinCmd = &cobra.Command{
	Use:   "join <address> <token>",
	Short: "Join a existing panel cluster",
	Long:  `Enable a new node to join an existing distributed cluster. Address can belong to any node in the cluster.`,
	Run:   joinAction,
}

func joinAction(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		_ = cmd.Help()
		return
	}

	address := args[0]
	token := args[1]

	ip, port := utils.AssertPublicAddress(address)

}

func init() {
	rootCmd.AddCommand(joinCmd)

	attachServerFlags(joinCmd)
}
