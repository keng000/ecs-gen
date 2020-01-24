package command

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/keng000/ecs-gen/skeleton"
	"github.com/keng000/ecs-gen/utils/logger"
)

const (
	envFile   string = ".ecs-gen.json"
	pjDirName string = "environments"
)

func searchEnv() (string, error) {
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

	envFilePath := filepath.Join(curPath, envFile)
	return envFilePath, nil
}

func existsIn(baseName, dstPath string) (bool, error) {
	files, err := ioutil.ReadDir(dstPath)
	if err != nil {
		logger.Errorf(err.Error())
		return false, err
	}

	for _, file := range files {
		if file.Name() == baseName {
			return true, nil
		}
	}

	return false, nil
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func loadExecutable(envFilePath string) (*skeleton.DumpExecutable, error) {
	content, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	dumpExecutable := &skeleton.DumpExecutable{}
	if err := json.Unmarshal(content, dumpExecutable); err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return dumpExecutable, nil
}

func writeExecutable(envFilePath string, data *skeleton.DumpExecutable) error {
	dumpData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	fp, err := os.Create(envFilePath)
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
