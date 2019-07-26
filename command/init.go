package command

import (
	"github.com/keng000/ecs-gen/skeleton"
	"github.com/urfave/cli"
)

// CmdInit process the init command
func CmdInit(c *cli.Context) error {
	executable := &skeleton.Executable{
		Project: c.String("project"),
		Region:  c.String("region"),
	}

	name := "environments"
	skeleton := skeleton.Skeleton{
		Path:       name,
		Executable: executable,
	}

	if err := skeleton.Base(); err != nil {
		// TODO: as logging
		println(err.Error())
	}

	return nil
}
