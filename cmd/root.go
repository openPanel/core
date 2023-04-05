package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name:                 "openPanel",
	Usage:                "Distributed linux panel",
	Suggest:              true,
	EnableBashCompletion: true,
	Commands: []*cli.Command{
		startCmd,
		joinCmd,
		resumeCmd,
	},
}

func Execute() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("%#v", err)
	}
}
