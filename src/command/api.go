package command

import (
	"github.com/keng000/ecs-gen/src/skeleton"
	"github.com/keng000/ecs-gen/src/utils/config"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func CmdAPI(c *cli.Context) error {
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

	if len(c.Args()) == 0 {
		log.Fatal("One or more api name should be passed to the args")
	}

	for _, apiName := range c.Args() {
		if contains(cfg.APIName, apiName) {
			log.WithFields(log.Fields{
				"api": apiName,
			}).Infof("API already exists. Do nothing")
			continue
		}

		s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
		if err := s.API(&skeleton.APIExecutable{
			Project: cfg.Project,
			APIName: apiName,
		}); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("Failed to Exec template")
		}

		cfg.APIName = append(cfg.APIName, apiName)
		log.WithFields(log.Fields{
			"api": apiName,
		}).Infof("API created")
	}

	if err := cfgCtrl.Write(cfg); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to dump config into file")
	}
	return nil
}
