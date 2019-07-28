package command

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/keng000/ecs-gen/skeleton"
	"github.com/urfave/cli"
)

func CmdAPI(c *cli.Context) error {
	envFilePath, err := searchEnv()
	if err != nil {
		log.Panic(err)
		return err
	}

	content, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		log.Panic(err)
		return err
	}

	dumpExecutable := &skeleton.DumpExecutable{}
	if err := json.Unmarshal(content, dumpExecutable); err != nil {
		log.Panic(err)
		return err
	}

	apiName := c.String("a")
	executable := &skeleton.Executable{
		Project: dumpExecutable.Project,
		Region:  dumpExecutable.Region,
		APIName: apiName,
	}

	path := filepath.Dir(envFilePath)
	skeleton := skeleton.Skeleton{
		Path:       path,
		Executable: executable,
	}

	if err := skeleton.API(); err != nil {
		log.Panic(err)
		return err
	}

	dumpExecutable.APIName = append(dumpExecutable.APIName, apiName)
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

	log.Printf("API created: %s", apiName)

	return nil
}
