package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/keng000/ecs-gen/utils/logger"
)

const (
	envFile   string = ".ecs-gen.json"
	pjDirName string = "environments"
)

type (
	// IConfigController is an interface for Controller
	IConfigController interface {
		Read() (*Config, error)
		Write(cfg *Config) error
	}

	// Controller read / write configs from file.
	Controller struct {
		ProjectRoot string
		ConfigPath  string
	}

	// Config contains the whole config which would dumped into secret file.
	Config struct {
		// Project is the name of the ecs project
		Project string

		// APIName is the name for auto scale generate
		APIName []string
	}
)

// Init creates a new config file
func Init() error {
	curPath, err := os.Getwd()
	if err != nil {
		return err
	}

	envDirExist, err := existsIn(pjDirName, curPath)
	if err != nil {
		return err
	}
	if envDirExist {
		return fmt.Errorf("%s directory already exists", pjDirName)
	}

	envFilePath := filepath.Join(curPath, pjDirName, envFile)
	if err := createEmpty(envFilePath); err != nil {
		return err
	}

	return nil
}

func createEmpty(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}

	emptyFile, err := os.Create(path)
	if err != nil {
		return err
	}
	emptyFile.Close()
	return nil
}

// NewController returns an instance which meets Controller
func NewController() (*Controller, error) {
	path, err := serchRoot()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	configFilePath := filepath.Join(path, envFile)
	return &Controller{
		ProjectRoot: path,
		ConfigPath:  configFilePath,
	}, nil
}

func (c *Controller) Read() (*Config, error) {
	content, err := ioutil.ReadFile(c.ConfigPath)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	cfg := &Config{}
	if err := json.Unmarshal(content, cfg); err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return cfg, nil
}

func (c *Controller) Write(cfg *Config) error {
	dumpData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	fp, err := os.Create(c.ConfigPath)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = fp.Write(dumpData)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

// serchRoot serch the project root path include config
func serchRoot() (string, error) {
	curPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	envDirExist, err := existsIn(pjDirName, curPath)
	if err != nil {
		return "", err
	}
	if envDirExist {
		curPath = filepath.Join(curPath, pjDirName)
	}

	for {
		exist, err := existsIn(envFile, curPath)
		if err != nil {
			return "", err
		}
		if exist {
			// .ecs-gen.json file found
			break
		}
		if curPath == "/" {
			return "", errors.New(".ecs-gen.json not found. please run ecs-gen init before")
		}

		curPath = filepath.Dir(curPath)
	}
	return curPath, nil
}

// existsIn checks if its with the `fileName` exists in the `dstPath` or not
func existsIn(fileName, dstPath string) (bool, error) {
	files, err := ioutil.ReadDir(dstPath)
	if err != nil {
		logger.Errorf(err.Error())
		return false, err
	}

	for _, file := range files {
		if file.Name() == fileName {
			return true, nil
		}
	}

	return false, nil
}
