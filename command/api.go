package command

import (
	"encoding/json"
	"fmt"
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
	dumpExecutable, err := loadExecutable(envFilePath)
	if err != nil {
		log.Panic(err)
		return err
	}
	apiName := c.String("n")
	if apiName == "" {
		log.Panic("-name, -n should be passed to the args")
		return fmt.Errorf("-name, -n should be passed to the args")
	} else if contains(dumpExecutable.APIName, apiName) {
		log.Panicf("already exists: %s", apiName)
		return fmt.Errorf("already exists: %s", apiName)
	}

	data := map[string]string{
		"APIname": apiName,
	}

	path := filepath.Dir(envFilePath)
	skeleton := skeleton.Skeleton{
		Path: path,
	}

	if err := skeleton.API(data); err != nil {
		log.Panic(err)
		return err
	}

	dumpExecutable.APIName = append(dumpExecutable.APIName, apiName)

	if err := writeExecutable(envFilePath, dumpExecutable); err != nil {
		log.Panic(err)
		return err
	}

	log.Printf("API created: %s", apiName)

	return nil
}

func loadExecutable(envFilePath string) (*skeleton.DumpExecutable, error) {
	content, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	dumpExecutable := &skeleton.DumpExecutable{}
	if err := json.Unmarshal(content, dumpExecutable); err != nil {
		log.Panic(err)
		return nil, err
	}
	return dumpExecutable, nil
}

func writeExecutable(envFilePath string, data *skeleton.DumpExecutable) error {
	dumpData, err := json.MarshalIndent(data, "", "  ")
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

	return nil
}
