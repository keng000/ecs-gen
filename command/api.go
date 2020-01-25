package command

import (
	"fmt"
	"log"

	"github.com/keng000/ecs-gen/skeleton"
	"github.com/keng000/ecs-gen/utils/config"
	"github.com/keng000/ecs-gen/utils/logger"
	"github.com/urfave/cli"
)

func CmdAPI(c *cli.Context) error {
	cfgCtrl, err := config.NewController()
	if err != nil {
		return err
	}

	cfg, err := cfgCtrl.Read()
	if err != nil {
		return err
	}

	if len(c.Args()) == 0 {
		logger.Error("One or more api name should be passed to the args")
		return fmt.Errorf("One or more api name should be passed to the args")
	}

	for _, apiName := range c.Args() {
		if contains(cfg.APIName, apiName) {
			log.Printf("Already exists: %s. Do nothing", apiName)
			continue
		}

		executable := skeleton.APIExecutable{
			Project: cfg.Project,
			APIName: apiName,
		}

		s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
		if err := s.API(executable); err != nil {
			logger.Error(err.Error())
			return err
		}

		cfg.APIName = append(cfg.APIName, apiName)
		log.Printf("API created: %s", apiName)
	}

	if err := cfgCtrl.Write(cfg); err != nil {
		return err
	}
	return nil
}
