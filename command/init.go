package command

import (
	"log"
	"os"
	"path/filepath"

	p_skeleton "github.com/keng000/ecs-gen/skeleton"
	"github.com/urfave/cli"
)

// CmdInit process the init command
func CmdInit(c *cli.Context) error {
	var project string
	if len(c.Args()) == 0 {
		log.Print("No project name defined. The project name will be `project`")
		project = "project"
	} else if len(c.Args()) > 1 {
		log.Printf("Too much arguments for init command. The following args will be ignored: %v", c.Args()[1:])
		project = c.Args().First()
	} else {
		project = c.Args().First()
	}

	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(curDir, pjDirName)
	_, err = os.Stat(path)
	if err == nil {
		log.Panicf("Cannot create directory. Project already exists: %s", path)
		// TODO: ask whether to force initialize
		return err
	}

	executable := p_skeleton.InitExecutable{
		Project: project,
	}

	skeleton := p_skeleton.Skeleton{
		Path: path,
	}

	if err := skeleton.Init(executable); err != nil {
		log.Panic(err)
		return err
	}

	envFilePath := filepath.Join(path, envFile)
	dumpExecutable := &p_skeleton.DumpExecutable{
		Project: executable.Project,
	}

	if err := writeExecutable(envFilePath, dumpExecutable); err != nil {
		log.Panic(err)
		return err
	}

	log.Print("environments created")
	return nil
}
