package main

import (
	"os"

	"github.com/keng000/ecs-gen/src/utils/logger"

	"github.com/urfave/cli"
)

func main() {

	logger.SetupLogger()

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "keng000"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
