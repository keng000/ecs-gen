package command

import (
	"github.com/keng000/ecs-gen/src/skeleton"
	"github.com/keng000/ecs-gen/src/utils/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func CmdDB(c *cli.Context) error {
	cfgCtrl, err := config.NewController()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Faild to create config controller")
	}

	if !cfgCtrl.PjAlreadyCreated {
		log.Fatal("No project found. run `ecs-gen init` before")
	}

	cfg, err := cfgCtrl.Read()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Faild to load config")
	}

	if c.NArg() == 0 {
		log.Fatal("No db name specified")
	} else if c.NArg() > 1 {
		log.WithFields(log.Fields{
			"db name": c.Args().Get(0),
		}).Warn("Multi db name specified. First one will use")
	}

	dbName := c.Args().Get(0)

	s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
	if err := s.DB(&skeleton.DBExecutable{Project: cfg.Project, DBName: dbName}); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to Exec template")
	}

	log.Info("DB created")

	return nil
}
