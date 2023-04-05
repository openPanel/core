package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/openPanel/core/app/bootstrap"
)

var resumeCmd = &cli.Command{
	Name:        "resume",
	Usage:       "Resume a paused node",
	Description: "Resume a pre-existing node that was paused.",
	Action: func(context *cli.Context) error {
		bootstrap.Resume()
		return nil
	},
}
