package command

import (
	"fmt"
	"log"

	"github.com/keng000/ecs-gen/src/skeleton"
	"github.com/keng000/ecs-gen/src/utils/config"
	"github.com/keng000/ecs-gen/src/utils/logger"

	mapset "github.com/deckarep/golang-set"
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

	cfg, err := cfgCtrl.Read()
	if err != nil {
		return err
	}

	if len(c.Args()) == 0 {
		log.Panic("One or more region name shold be passed to the deploy command")
	}

	regionSet := mapset.NewSetFromSlice(toInterfaceSlice(cfg.Region))
	for _, region := range c.Args() {
		if !isValid(region) {
			logger.Error("Invalid region name: %s\n", region)
			continue
		}
		regionSet.Add(region)
	}

	for region := range regionSet.Iterator().C {
		region := region.(string)
		if !isValid(region) {
			logger.Error("Invalid region name: %s\n", region)
			continue
		}

		s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
		if err := s.Deploy(&skeleton.DeployExecutable{
			Project: cfg.Project,
			Region:  region,
			APIName: cfg.APIName,
		}); err != nil {
			logger.Error("Failed to Exec template")
			logger.Error(err.Error())
			return err
		}

		logger.Infof("Region updated: %s", region)
	}

	cfg.Region = toStringSlice(regionSet.ToSlice())
	if err := cfgCtrl.Write(cfg); err != nil {
		return err
	}
	return nil
}
