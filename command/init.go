package command

import (
	"log"

	"github.com/keng000/ecs-gen/skeleton"
	"github.com/keng000/ecs-gen/utils/config"
	"github.com/keng000/ecs-gen/utils/logger"

	"github.com/urfave/cli"
)

// CmdInit process the init command
func CmdInit(c *cli.Context) error {
	project := c.String("project")
	log.Printf("[INFO] project initialized with name `%s`\n", project)

	if err := config.Init(); err != nil {
		return err
	}

	cfgCtrl, err := config.NewController()
	if err != nil {
		return err
	}

	s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
	if err := s.Init(skeleton.InitExecutable{Project: project}); err != nil {
		logger.Error(err.Error())
		return err
	}

	cfg := &config.Config{Project: project}
	if err := cfgCtrl.Write(cfg); err != nil {
		return err
	}

	log.Print("environments created")
	return nil
}
