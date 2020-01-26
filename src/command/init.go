package command

import (
	"github.com/keng000/ecs-gen/src/skeleton"
	"github.com/keng000/ecs-gen/src/utils/config"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// CmdInit process the init command
func CmdInit(c *cli.Context) error {
	if c.NArg() == 0 {
		log.Fatal("No project name specified")
	} else if c.NArg() > 1 {
		log.WithFields(log.Fields{
			"project": c.Args().Get(0),
		}).Warn("Multi project name specified. First one will use")
	}

	project := c.Args().Get(0)
	cfgCtrl, err := config.NewController()
	if err != nil {
		log.Fatal(err.Error())
	}

	if cfgCtrl.PjAlreadyCreated {
		log.WithFields(log.Fields{
			"path": cfgCtrl.ProjectRoot,
		}).Fatal("envrionment already exists")
	}

	if err := cfgCtrl.Init(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to init config file")
	}

	s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
	if err := s.Init(&skeleton.InitExecutable{Project: project}); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to Exec template")
	}

	cfg := &config.Config{Project: project}
	if err := cfgCtrl.Write(cfg); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to dump config into file")
	}

	log.WithFields(log.Fields{
		"name": project,
	}).Infof("Project initialized. Environments created")
	return nil
}
