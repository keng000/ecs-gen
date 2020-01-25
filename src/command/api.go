package command

import (
	"github.com/keng000/ecs-gen/src/skeleton"
	"github.com/keng000/ecs-gen/src/utils/config"
	"github.com/keng000/ecs-gen/src/utils/logger"

	"github.com/urfave/cli"
)

func CmdAPI(c *cli.Context) error {
	cfgCtrl, err := config.NewController()
	if err != nil {
		logger.Error("Faild to create config controller")
		panic(err)
	}

	if !cfgCtrl.PjAlreadyCreated {
		msg := "No project found. run `ecs-gen init` before"
		logger.Error(msg)
		panic(msg)
	}

	cfg, err := cfgCtrl.Read()
	if err != nil {
		logger.Error("Faild to load config")
		panic(err)
	}

	if len(c.Args()) == 0 {
		msg := "One or more api name should be passed to the args"
		logger.Error(msg)
		panic(msg)
	}

	for _, apiName := range c.Args() {
		if contains(cfg.APIName, apiName) {
			logger.Infof("Already exists: %s. Do nothing", apiName)
			continue
		}

		s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
		if err := s.API(&skeleton.APIExecutable{
			Project: cfg.Project,
			APIName: apiName,
		}); err != nil {
			logger.Error("Failed to Exec template")
			logger.Error(err.Error())
			panic(err)
		}

		cfg.APIName = append(cfg.APIName, apiName)
		logger.Infof("API created: %s", apiName)
	}

	if err := cfgCtrl.Write(cfg); err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	return nil
}
