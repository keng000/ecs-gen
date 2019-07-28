package command

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	envFile   string = ".ecs-gen.json"
	pjDirName string = "environment"
)

func searchEnv() (string, error) {
	curPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	curPath = filepath.Join(curPath, pjDirName)

	for {
		exist, err := exists(envFile, curPath)
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
	if err != nil {
		panic(err)
	}
	return envFilePath, nil
}

func exists(filename, dstPath string) (bool, error) {
	files, err := ioutil.ReadDir(dstPath)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if file.Name() == filename {
			return true, nil
		}
	}

	return false, nil
}
