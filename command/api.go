package command

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/keng000/ecs-gen/skeleton"
	"github.com/keng000/ecs-gen/utils/logger"
	"github.com/urfave/cli"
)

func CmdAPI(c *cli.Context) error {
	envFilePath, err := searchEnv()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	dumpExecutable, err := loadExecutable(envFilePath)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if len(c.Args()) == 0 {
		logger.Error("One or more api name should be passed to the args")
		return fmt.Errorf("One or more api name should be passed to the args")
	}

	for _, apiName := range c.Args() {
		if contains(dumpExecutable.APIName, apiName) {
			log.Printf("Already exists: %s. Do nothing", apiName)
			continue
		}

		executable := skeleton.APIExecutable{
			Project: dumpExecutable.Project,
			APIName: apiName,
		}

		path := filepath.Dir(envFilePath)

		s := skeleton.Skeleton{Path: path}
		if err := s.API(executable); err != nil {
			logger.Error(err.Error())
			return err
		}

		dumpExecutable.APIName = append(dumpExecutable.APIName, apiName)
		log.Printf("API created: %s", apiName)
	}

	if err := writeExecutable(envFilePath, dumpExecutable); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
