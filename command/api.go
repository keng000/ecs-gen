package command

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/keng000/ecs-gen/skeleton"
	p_skeleton "github.com/keng000/ecs-gen/skeleton"
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

	if len(c.Args()) == 0 {
		log.Panic("One or more api name should be passed to the args")
		return fmt.Errorf("One or more api name should be passed to the args")
	}

	for _, apiName := range c.Args() {
		if contains(dumpExecutable.APIName, apiName) {
			log.Printf("Already exists: %s. Do nothing", apiName)
			continue
		}

		executable := p_skeleton.APIExecutable{
			Project: dumpExecutable.Project,
			APIName: apiName,
		}

		path := filepath.Dir(envFilePath)
		skeleton := skeleton.Skeleton{
			Path: path,
		}

		if err := skeleton.API(executable); err != nil {
			log.Panic(err)
			return err
		}

		dumpExecutable.APIName = append(dumpExecutable.APIName, apiName)
		log.Printf("API created: %s", apiName)
	}

	if err := writeExecutable(envFilePath, dumpExecutable); err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
