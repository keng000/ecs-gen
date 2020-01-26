package command

import (
	"errors"
	"fmt"

	"github.com/keng000/ecs-gen/src/skeleton"
	"github.com/keng000/ecs-gen/src/utils/config"
	"github.com/keng000/ecs-gen/src/utils/logger"

	"github.com/urfave/cli"
)

// CmdInit process the init command
func CmdInit(c *cli.Context) error {
	if c.NArg() == 0 {
		logger.Error("No project name specified")
		return errors.New("No project name specified")
	} else if c.NArg() > 1 {
		logger.Warn("Multi project name specified. First one will use")
	}

	project := c.Args().Get(0)
	cfgCtrl, err := config.NewController()
	if err != nil {
		return err
	}

	if cfgCtrl.PjAlreadyCreated {
		msg := fmt.Sprintf("envrionment already exists: %s", cfgCtrl.ProjectRoot)
		logger.Error(msg)
		return errors.New(msg)
	}

	if err := cfgCtrl.Init(); err != nil {
		logger.Error("Failed to init config file")
		return err
	}

	s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
	if err := s.Init(&skeleton.InitExecutable{Project: project}); err != nil {
		logger.Error("Failed to Exec template")
		logger.Error(err.Error())
		return err
	}

	cfg := &config.Config{Project: project}
	if err := cfgCtrl.Write(cfg); err != nil {
		logger.Error("Failed to dump config into file")
		logger.Error(err.Error())
		return err
	}

	logger.Infof("Project initialized with name `%s`\n", project)
	logger.Info("Environments created")
	return nil
}
