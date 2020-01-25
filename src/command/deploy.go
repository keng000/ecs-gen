package command

import (
	"fmt"
	"log"

	"github.com/keng000/ecs-gen/src/skeleton"
	"github.com/keng000/ecs-gen/src/utils/config"
	"github.com/keng000/ecs-gen/src/utils/logger"

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

	for _, region := range c.Args() {
		if !isValid(region) {
			logger.Error("Invalid region name: %s\n", region)
			continue
		}
		if contains(cfg.Region, region) {
			logger.Infof("Already exists: %s. Do nothing", region)
			continue
		}

		executable := skeleton.DeployExecutable{
			Project: cfg.Project,
			Region:  region,
		}

		s := skeleton.Skeleton{Path: cfgCtrl.ProjectRoot}
		if err := s.Deploy(executable); err != nil {
			logger.Error(err.Error())
			return err
		}

		cfg.Region = append(cfg.Region, region)
		logger.Infof("Region created: %s", region)
	}

	if err := cfgCtrl.Write(cfg); err != nil {
		return err
	}
	return nil
}

func isValid(region string) bool {
	return regions.Contains(region)
}
