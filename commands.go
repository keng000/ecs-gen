package main

import (
	"fmt"
	"os"

	"github.com/keng000/ecs-gen/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "ENV_A",
		Name:   "a",
		Value:  "",
		Usage:  "",
	},
	cli.StringFlag{
		EnvVar: "ENV_B",
		Name:   "b",
		Value:  "",
		Usage:  "",
	},
}

var Commands = []cli.Command{
	{
		Name:   "init",
		Usage:  "Define basic infrastructure. e.g. VPC, Subnet, SG, TG...",
		Action: command.CmdInit,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project, p",
				Value: "sample-pj",
				Usage: "",
			},
			cli.StringFlag{
				Name:  "region, r",
				Value: "ap-northeast-1",
				Usage: "",
			},
		},
	},
	{
		Name:   "auto-scale",
		Usage:  "",
		Action: command.CmdAutoScale,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
