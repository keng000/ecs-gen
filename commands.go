package main

import (
	"fmt"
	"os"

	"github.com/keng000/ecs-gen/src/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:      "init",
		Usage:     "Define basic infrastructure. e.g. VPC, Subnet, SG, TG...",
		ArgsUsage: "[PROJECT]",
		Action:    command.CmdInit,
	},
	{
		Name:      "api",
		Usage:     "Generate target groups and ecr repos with auto scale setting",
		ArgsUsage: "[APIs...]",
		Action:    command.CmdAPI,
	},
	{
		Name:      "deploy",
		Usage:     "Generate region directory for module deploy",
		ArgsUsage: "[REGIONs...]",
		Action:    command.CmdDeploy,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "list, l",
				Usage: "List available region list",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
