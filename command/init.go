package command

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	p_skeleton "github.com/keng000/ecs-gen/skeleton"
	"github.com/urfave/cli"
)

// CmdInit process the init command
func CmdInit(c *cli.Context) error {
	executable := &p_skeleton.Executable{
		Project: c.String("project"),
		Region:  c.String("region"),
	}

	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(curDir, pjDirName)
	_, err = os.Stat(path)
	if err == nil {
		log.Panicf("cannot create directory. already exists: %s", path)
		return err
	}

	skeleton := p_skeleton.Skeleton{
		Path:       path,
		Executable: executable,
	}

	if err := skeleton.Init(); err != nil {
		log.Panic(err)
		return err
	}

	envFilePath := filepath.Join(path, envFile)
	dumpExecutable := &p_skeleton.DumpExecutable{
		Project: executable.Project,
		Region:  executable.Region,
	}

	dumpData, err := json.MarshalIndent(dumpExecutable, "", "  ")
	if err != nil {
		log.Panic(err)
		return err
	}

	fp, err := os.Create(envFilePath)
	if err != nil {
		log.Panic(err)
		return err
	}

	_, err = fp.Write(dumpData)
	if err != nil {
		log.Panic(err)
		return err
	}

	log.Print("environments created")
	return nil
}
