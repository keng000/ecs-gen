package command

import (
	"fmt"

	"github.com/keng000/ecs-gen/src/skeleton"
	"github.com/keng000/ecs-gen/src/utils/config"

	mapset "github.com/deckarep/golang-set"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// CmdDeploy will generate a directory for a specific region to deploy environments
func CmdDeploy(c *cli.Context) error {
	// Print out available retgion list
	if c.Bool("list") {
		fmt.Println("*** Available region list ***")
		for _, s := range regions.AsSlice() {
			fmt.Printf("  - %s\n", s)
		}
		return nil
	}

	// Main
	cfgCtrl, err := config.NewController()
	if err != nil {
		return err
	}

	if !cfgCtrl.PjAlreadyCreated {
		log.Fatal("No project found. run `ecs-gen init` before")
	}

	cfg, err := cfgCtrl.Read()
	if err != nil {
		return err
	}

	if len(c.Args()) == 0 {
		log.Fatal("One or more region name shold be passed to the deploy command")
	}

	regionSet := mapset.NewSetFromSlice(toInterfaceSlice(cfg.Region))
	for _, region := range c.Args() {
		if !isValid(region) {
			log.WithFields(log.Fields{
				"region": region,
			}).Error("Invalid region name")
			continue
		}
		regionSet.Add(region)
	}

	for region := range regionSet.Iterator().C {
		region := region.(string)

		s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
		if err := s.Deploy(&skeleton.DeployExecutable{
			Project: cfg.Project,
			Region:  region,
			APIName: cfg.APIName,
		}); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("Failed to Exec template")
		}

		log.WithFields(log.Fields{
			"region": region,
		}).Infof("Region updated")
	}

	cfg.Region = toStringSlice(regionSet.ToSlice())
	if err := cfgCtrl.Write(cfg); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to dump config into file")
	}
	return nil
}
