package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openPanel",
	Short: "Distributed linux panel",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("%#v", err)
	}
}
