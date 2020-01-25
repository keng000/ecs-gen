package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	envFile   string = ".ecs-gen.json"
	pjDirName string = "environments"
)

type (
	// IConfigController is an interface for Controller
	IConfigController interface {
		Init() error
		Read() (*Config, error)
		Write(cfg *Config) error
	}

	// Controller read / write configs from file.
	Controller struct {
		ProjectRoot      string
		ConfigPath       string
		PjAlreadyCreated bool
	}

	// Config contains the whole config which would dumped into secret file.
	Config struct {
		// Project is the name of the ecs project
		Project string

		// APIName is the name for auto scale generate
		APIName []string

		Region []string
	}
)

// NewController returns an instance which meets Controller
func NewController() (*Controller, error) {
	c := &Controller{}
	path, err := searchRoot()
	if path == "/" {
		c.PjAlreadyCreated = false
		return c, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to search root")
	}
	c.ProjectRoot = path
	c.ConfigPath = filepath.Join(path, envFile)
	c.PjAlreadyCreated = true
	return c, nil
}

// // Search search and store the environment project dir path and config file path
// func (c *Controller) Search() error {

// 	c.ProjectRoot = path
// 	c.ConfigPath = filepath.Join(path, envFile)
// 	return nil
// }

// Read loads config
func (c *Controller) Read() (*Config, error) {
	content, err := ioutil.ReadFile(c.ConfigPath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := json.Unmarshal(content, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Write dums config
func (c *Controller) Write(cfg *Config) error {
	dumpData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	fp, err := os.Create(c.ConfigPath)
	if err != nil {
		return err
	}

	_, err = fp.Write(dumpData)
	if err != nil {
		return err
	}

	return nil
}

// Init creates a new config file
func (c *Controller) Init() error {
	curPath, err := os.Getwd()
	if err != nil {
		return err
	}

	pjDir := filepath.Join(curPath, pjDirName)
	if _, err := os.Stat(pjDir); !os.IsNotExist(err) {
		return fmt.Errorf("%s directory already exists", pjDirName)
	}

	envFilePath := filepath.Join(curPath, pjDirName, envFile)
	if err := createEmpty(envFilePath); err != nil {
		return err
	}

	c.ProjectRoot = pjDir
	c.ConfigPath = envFilePath
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

// searchRoot search the project root path include config
func searchRoot() (string, error) {
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
			return "/", errors.New(".ecs-gen.json not found")
		}

		curPath = filepath.Dir(curPath)
	}
	return curPath, nil
}

// existsIn checks if its with the `fileName` exists in the `dstPath` or not
func existsIn(fileName, dstPath string) (bool, error) {
	pjDir := filepath.Join(dstPath, fileName)
	if _, err := os.Stat(pjDir); !os.IsNotExist(err) {
		return true, nil
	}
	return false, nil
}
